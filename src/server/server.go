package server

import (
	"os"
	"url_manager/database"
	"url_manager/middleware"
	"url_manager/repositories"
	"url_manager/session"

	"url_manager/api"

	"github.com/gin-gonic/gin"
)

func Open(port string) {
	router := gin.Default()

	// initialize session.
	router.Use(session.NewSessionHandler())

	// initialize favicon.
	router.Static("favicon.ico", "app/assets/favicon.ico")
	router.Use(middleware.NewFaviconHandler())

	// initialize google authentication.
	redirectUrl := os.Getenv("REDIRECT_URL")
	credFilePath := os.Getenv("CRED_FILE_PATH")
	secret := []byte(os.Getenv("GOOGLE_SECRET"))
	scopes := []string{"openid"}
	middleware.Setup(redirectUrl, credFilePath, scopes, secret)
	router.Use(middleware.NewAuthHandler())

	// initialize main.
	{
		api := api.NewLinksApi(
			repositories.NewUserRepository(database.Database()),
			repositories.NewLinkRepository(database.Database()),
			repositories.NewLinkListRepository(database.Database()),
			repositories.NewLinkListRelationRepository(database.Database()),
		)
		users := router.Group("/users/:user_id/links")
		users.GET("", api.Index)
		users.GET("/:link_id", api.Show)
		users.POST("", api.Create)
		users.PUT("", api.Update)
		users.DELETE("/:link_id", api.Delete)
	}

	router.Run(port)
}
