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

	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("32bytes-secret-auth-key"))
	if err != nil {
		panic(err)
	}
	router.Use(sessions.Sessions("URLManager", store))

	router.POST("/sign_in", controllers.CreateSession)
	router.DELETE("/sign_out", controllers.DestroySession)

	users := router.Group("/users")
	{
		users.GET("", controllers.ShowUsers)
		users.POST("", controllers.CreateUser)
		users.GET("/:id", middlewares.RequireLogin(), controllers.ShowUser)
		users.PUT("/:id", middlewares.RequireLogin(), controllers.UpdateUser)
		users.DELETE("/:id", middlewares.RequireLogin(), controllers.DeleteUser)
	}

	router.Run(":8080")
}
