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

	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("32bytes-secret-auth-key"))
	if err != nil {
		panic(err)
	}
	router.Use(sessions.Sessions("URLManager", store))

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

	router.GET("/", controllers.About)
	router.GET("/about", controllers.About)
	router.GET("/contact", controllers.Contact)

	router.GET("/login", controllers.NewSession)
	router.POST("/login", controllers.CreateSession)
	router.GET("/logout", controllers.DestroySession)

	users := router.Group("/users")
	{
		// users.GET("", controllers.ShowUsers)
		users.GET("/:id", middlewares.RequireLogin(), controllers.ShowUser)
		users.GET("/new", controllers.NewUser)
		users.POST("", controllers.CreateUser)
		// users.GET("/edit", controllers.EditURL)
		// users.PUT("/:id", middlewares.RequireLogin(), controllers.UpdateUser)
		// users.DELETE("/:id", middlewares.RequireLogin(), controllers.DeleteUser)

	}

	urls := router.Group("/users/:id/urls")
	{
		// urls.GET("", controllers.ShowURLs)
		// urls.GET("/:id", controllers.ShowURL)
		urls.GET("/new", controllers.NewURL)
		urls.POST("", controllers.CreateURL)
		// urls.GET("/:url_id/edit", controllers.EditURL)
		// urls.PUT("/:url_id", controllers.UpdateURL)
		urls.GET("/:urlID/delete", controllers.DeleteURL)
	}

	router.Run(":8080")
}
