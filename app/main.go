package main

import (
	"url_manager/controllers"
	"url_manager/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	router := gin.Default()

	{
		auth_controller := controllers.AuthController{}
		router.POST("/sign_in", auth_controller.SignIn)
		// router.POST("/sign_up", auth_controller.SignOut)
	}

	users := router.Group("/users")
	{
		user_controller := controllers.UserController{}
		// users.GET("", user_controller.Index)
		users.POST("", user_controller.Create)
		// users.GET("/:id", user_controller.Show)
		// users.PUT("/:id", user_controller.Update)
		// users.DELETE("/:id", user_controller.Delete)
	}

	router.Run(":8080")
}
