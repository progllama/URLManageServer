package server

import (
	"net/http"
	"os"
	"url_manager/database"
	"url_manager/middleware"
	"url_manager/repositories"
	"url_manager/session"

	"url_manager/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Open(port string) {
	router := gin.Default()

	// load template.
	router.LoadHTMLGlob("template/*")

	// initialize session.
	router.Use(session.NewSessionHandler())

	// initialize favicon.
	// router.Static("favicon.ico", "app/assets/favicon.ico")
	// router.Use(middleware.NewFaviconHandler())

	// initialize static
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	//initialize home
	router.GET("/home", func(ctx *gin.Context) {
		stateValue := middleware.RandToken()
		session := sessions.Default(ctx)
		session.Set("state", stateValue)
		session.Save()

		data := []HomeData{
			{"title1", "/hogehoge"},
			{"title2", "/hogehoge"},
			{"title3", "/hogehoge"},
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"auth_url":  middleware.GetLoginURL(stateValue),
			"logged_in": session.Get("logged_in"),
			"data":      data,
		})
	})

	// initialize google authentication.
	redirectUrl := os.Getenv("REDIRECT_URL")
	credFilePath := os.Getenv("CRED_FILE_PATH")
	secret := []byte(os.Getenv("GOOGLE_SECRET"))
	scopes := []string{"openid"}
	middleware.SetUserRepository(repositories.NewUserRepository(database.Database()))
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

type HomeData struct {
	Title string
	Url   string
}
