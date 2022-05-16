package server

import (
	"url_manager/app/controllers"
	"url_manager/app/middlewares"
	"url_manager/app/repositories"
	"url_manager/app/session"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

// settings
var (
	sessionFactory = session.NewMemSessionFactory()
)

func Open(port string) {
	router := gin.Default()

	router.LoadHTMLGlob("app/templates/**/*")
	router.Static("/css", "app/assets/css")
	router.Static("/js", "app/assets/js")
	router.Static("favicon.ico", "app/assets/favicon.ico")

	// store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("32bytes-secret-auth-key"))
	// if err != nil {
	// 	panic(err)
	// }
	// router.Use(sessions.Sessions("URLManager", store))

	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(middlewares.ServeFavicon("./favicon.ico"))

	// router.Use(cors.New(cors.Config{
	// 	// 許可したいHTTPメソッドの一覧
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"OPTIONS",
	// 		"PUT",
	// 		"DELETE",
	// 	},
	// 	// 許可したいHTTPリクエストヘッダの一覧
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Headers",
	// 		"Content-Type",
	// 		"Content-Length",
	// 		"Accept-Encoding",
	// 		"X-CSRF-Token",
	// 		"Authorization",
	// 	},
	// 	// 許可したいアクセス元の一覧
	// 	AllowOrigins: []string{
	// 		"http://localhost:80",
	// 	},
	// 	AllowCredentials: true,
	// 	// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	//  return origin == "https://www.example.com:8080"
	// 	// },
	// 	// preflight requestで許可した後の接続可能時間
	// 	// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
	// 	MaxAge: 24 * time.Hour,
	// }))

	{
		ctrl := controllers.NewSessionController(repositories.GormNewUserRepository())
		router.GET("/login", ctrl.NewSession)
		router.POST("/login", ctrl.CreateSession)
		router.DELETE("/logout", ctrl.DestroySession)
	}

	users := router.Group("/users")
	{
		ctrl := controllers.NewUserController(repositories.GormNewUserRepository(), sessionFactory)
		users.GET("", ctrl.ShowAll)
		users.GET("/:id", ctrl.Show)
		users.GET("/new", ctrl.New)
		users.POST("", ctrl.Create)
		users.GET("/:id/edit", middlewares.RequireLogin(), ctrl.Edit)
		users.PUT("/:id", middlewares.RequireLogin(), ctrl.Update)
		users.PATCH("/:id", middlewares.RequireLogin(), ctrl.Update)
		users.DELETE("/:id", middlewares.RequireLogin(), ctrl.Delete)

		router.GET("/", middlewares.RequireLogin(), ctrl.Show)

		urls := users.Group("/:id/urls")
		{
			ctrl := controllers.NewUrlsController(repositories.GormNewUserRepository())
			urls.GET("", ctrl.ShowURLs)
			urls.GET("/:url_id", ctrl.ShowURL)
			urls.GET("/new", middlewares.RequireLogin(), ctrl.NewURL)
			urls.POST("", middlewares.RequireLogin(), ctrl.CreateURL)
			urls.GET("/:url_id/edit", middlewares.RequireLogin(), ctrl.EditURL)
			urls.PUT("/:url_id", middlewares.RequireLogin(), ctrl.UpdateURL)
			urls.DELETE("/:url_id", middlewares.RequireLogin(), ctrl.DeleteURL)
		}
	}

	router.Run(":8000")
}
