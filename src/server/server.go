package server

import (
	"url_manager/app/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Open(port string) {
	router := gin.Default()

	// TODO エラーハンドリング追加
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("32bytes-secret-auth-key"))
	router.Use(sessions.Sessions("URLManager", store))

	// auth_controller := controllers.AuthController{}
	// router.GET("/sign_in", auth_controller.SignIn)
	// router.POST("/sign_in", CreateSession)
	// router.POST("/sign_out", DestroySession)
	// router.GET("/sign_up", auth_controller.SignUp)

	users := router.Group("/users")
	{
		user_controller := controllers.UserController{}
		// auth_controller := controllers.AuthController{}

		// users.Use(auth_controller.SessionCheck)

		users.GET("", user_controller.Index)
		users.POST("", user_controller.Create)
		users.GET("/:id", user_controller.Show)
		// users.PUT("/:id", user_controller.Update)
		// users.DELETE("/:id", user_controller.Delete)
	}

	router.Run(":8080")
}
