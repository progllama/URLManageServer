package server

import (
	"os"
	"url_manager/database"
	middleware "url_manager/middlewares"
	"url_manager/repositories"
	"url_manager/session"

	"url_manager/api"

	"github.com/gin-gonic/gin"
)

func Open(port string) {
	router := gin.Default()

	session.InitializeSession(router)

	router.Static("favicon.ico", "app/assets/favicon.ico")
	router.Use(middleware.ServeFavicon("./favicon.ico"))

	redirectUrl := os.Getenv("REDIRECT_URL")
	credFilePath := os.Getenv("CRED_FILE_PATH")
	secret := []byte(os.Getenv("GOOGLE_SECRET"))
	scopes := []string{"https://www.googleapis.com/auth/userinfo.email", "openid"}
	middleware.Setup(redirectUrl, credFilePath, scopes, secret)
	router.Use(middleware.Auth())

	{
		api := api.NewLinksApi(
			repositories.NewUserRepository(database.Database()),
			repositories.NewLinkRepository(database.Database()),
			repositories.NewLinkListRepository(database.Database()),
			repositories.NewLinkListRelationRepository(database.Database()),
		)
		router4user := router.Group("/users/:user_id/links")
		router4user.GET("", api.Index)
		router4user.GET("/:link_id", api.Show)
		router4user.POST("", api.Create)
		router4user.PUT("", api.Update)
		router4user.DELETE("/:link_id", api.Delete)
	}

	router.Run(port)
}
