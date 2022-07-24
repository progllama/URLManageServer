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
	goauth "google.golang.org/api/oauth2/v2"
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

	router.GET("/home", func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		var loggedIn = false

		var userInfo goauth.Userinfo
		if v := session.Get("ginoauth_google_session"); v != nil {
			userInfo = v.(goauth.Userinfo)
			loggedIn = true
		} else {
			userInfo = goauth.Userinfo{}
		}

		var loginUrl = ""
		if !loggedIn {
			stateValue := middleware.RandToken()
			session.Set("state", stateValue)
			session.Save()
			loginUrl = middleware.GetLoginURL(stateValue)
		}

		repo := repositories.NewUserRepository(database.Database())
		user, err := repo.FindByOpenID(userInfo.Id)
		if err != nil {
			// ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		listRepo := repositories.NewLinkListRepository(database.Database())
		lists, err := listRepo.FindByUserId(int(user.ID))

		data := []HomeData{}
		for _, v := range lists {
			data = append(data, HomeData{
				ID:    int(v.ID),
				Title: v.Title,
			})
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"user_id":   user.ID,
			"auth_url":  loginUrl,
			"logged_in": loggedIn,
			"data":      data,
		})
	})

	router.GET("/auth/logout", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
		ctx.Redirect(302, "/home")
	})

	// initialize google authentication.
	redirectUrl := os.Getenv("REDIRECT_URL")
	credFilePath := os.Getenv("CRED_FILE_PATH")
	secret := []byte(os.Getenv("GOOGLE_SECRET"))
	scopes := []string{"openid"}
	middleware.SetUserRepository(repositories.NewUserRepository(database.Database()))
	middleware.Setup(redirectUrl, credFilePath, scopes, secret)
	router.Use(middleware.NewAuthHandler())

	router.GET("/auth/google", api.SignIn(repositories.NewUserRepository(database.Database())))
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

	{
		api := api.NewLinkListsApi(repositories.NewLinkListRepository(database.Database()))
		router.POST("/users/:user_id/link_lists", api.Create)
		router.DELETE("/users/:user_id/link_lists/:list_id", api.Delete)
	}

	router.Run(port)
}

type HomeData struct {
	ID    int
	Title string
	Url   string
}
