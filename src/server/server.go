package server

import (
	"net/http"
	"os"
	"url_manager/database"
	"url_manager/middleware"
	"url_manager/repositories"
	"url_manager/session"

	"url_manager/api"

	"github.com/gin-gonic/gin"
	goauth "google.golang.org/api/oauth2/v2"
)

func Open(port string) {
	router := gin.Default()

	// initialize session.
	router.Use(session.NewSessionHandler())

	if os.Getenv("MODE") == "dev" {
		router.GET("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "At least app is running.")
		})
		router.GET("/test/login", middleware.RawLoginHandler)
	}

	// initialize favicon.
	// router.Static("favicon.ico", "app/assets/favicon.ico")
	// router.Use(middleware.NewFaviconHandler())

	// initialize google authentication.
	redirectUrl := os.Getenv("REDIRECT_URL")
	credFilePath := os.Getenv("CRED_FILE_PATH")
	secret := []byte(os.Getenv("GOOGLE_SECRET"))
	scopes := []string{"openid"}
	middleware.SetUserRepository(repositories.NewUserRepository(database.Database()))
	middleware.Setup(redirectUrl, credFilePath, scopes, secret)
	router.Use(middleware.NewAuthHandler())
	if os.Getenv("MODE") == "dev" {
		router.GET("/auth/google", func(ctx *gin.Context) {
			var (
				res goauth.Userinfo
				ok  bool
			)

			val := ctx.MustGet("user")
			if res, ok = val.(goauth.Userinfo); !ok {
				res = goauth.Userinfo{Name: "no user"}
			}
			ctx.String(http.StatusOK, "success login"+res.Id)
		})
	}

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
