package server

import (
	"time"
	"url_manager/app/controllers"
	"url_manager/app/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Open(port string) {
	router := gin.Default()

	router.LoadHTMLGlob("app/templates/**/*")
	router.Static("/css", "app/assets/css")
	router.Static("/js", "app/assets/js")
	router.Static("favcion.ico", "app/assets/favicon.ico")

	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("32bytes-secret-auth-key"))
	if err != nil {
		panic(err)
	}
	router.Use(sessions.Sessions("URLManager", store))

	router.Use(middlewares.ServeFavicon("./favicon.ico"))

	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
		// AllowOriginFunc: func(origin string) bool {
		//  return origin == "https://www.example.com:8080"
		// },
		// preflight requestで許可した後の接続可能時間
		// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
		MaxAge: 24 * time.Hour,
	}))

	{
		ctrl := controllers.NewSessionController()
		router.GET("/login", ctrl.NewSession)
		router.POST("/login", ctrl.CreateSession)
		router.DELETE("/logout", ctrl.DestroySession)
	}

	users := router.Group("/users")
	{
		ctrl := controllers.NewUserController()
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
			urls.GET("", controllers.ShowURLs)
			urls.GET("/:url_id", controllers.ShowURL)
			urls.GET("/new", middlewares.RequireLogin(), controllers.NewURL)
			urls.POST("", middlewares.RequireLogin(), controllers.CreateURL)
			urls.GET("/:url_id/edit", middlewares.RequireLogin(), controllers.EditURL)
			urls.PUT("/:url_id", middlewares.RequireLogin(), controllers.UpdateURL)
			urls.DELETE("/:url_id", middlewares.RequireLogin(), controllers.DeleteURL)
		}
	}

	router.Run(":8000")
}
