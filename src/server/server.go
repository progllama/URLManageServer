package server

import (
	"url_manager/app/controllers"
	"url_manager/app/middlewares"

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

	router.GET("/", controllers.About)
	router.GET("/about", controllers.About)
	router.GET("/contact", controllers.Contact)

	router.GET("/login", controllers.NewSession)
	router.POST("/login", controllers.CreateSession)
	router.DELETE("/logout", controllers.DestroySession)

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
		// urls.GET("/edit", controllers.EditURL)
		// urls.PUT("/:id", controllers.UpdateURL)
		urls.DELETE("/:id", controllers.DeleteURL)
	}

	router.Run(":8080")
}
