package main

import (
	"url_manager/controllers"
	"url_manager/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	{
		auth_controller := AuthController{}
		router.Use(auth_controller.SessionCheck)
		router.GET("/sign_in", auth_controller.SignIn)
		router.POST("/sing_in", auth_controller.CreateSession)
		router.POST("/sign_out", auth_controller.DestroySession)
		router.GET("/sign_up", auth_controller.SignUp)
	}

	users := router.Group("/users")
	{
		user_controller := controllers.UserController{}

		users.Use(auth_controller.SessionCheck)

		users.GET("", user_controller.Index)
		users.POST("", user_controller.Create)
		users.GET("/:id", user_controller.Show)
		// users.PUT("/:id", user_controller.Update)
		// users.DELETE("/:id", user_controller.Delete)
	}

	router.Run(":8080")
}
