package app

import (
	"url_manager/controller"
	"url_manager/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Run() {
	app := &Application{}
	app.Run()
}

type Application struct {
}

func (app *Application) Run() {
	e := gin.Default()

	// Use session.
	store := memstore.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("url-plumber", store))

	// Use templates.
	loadTemplates(TEMPLATES)
	e.HTMLRender = Renderer

	// Use Statics.
	e.Static("/asset/js", JSFILES)
	e.Static("/asset/css", CSSFILES)

	// Static.
	static := e.Group("/")
	static.Use(middleware.Auth())
	static.GET("/", controller.Home)
	static.GET("/home", controller.Home)

	// Auth.
	auth := e.Group("/auth")
	auth.Use(middleware.Auth())
	auth.GET("/sign_up", controller.SignUpPage)
	auth.POST("/sign_up", controller.SignUp)
	auth.GET("/login", controller.LoginPage)
	auth.POST("/login", controller.Login)
	auth.GET("/logout", controller.Logout)

	users := e.Group("/users/:loginId")
	// Links.
	links := users.Group("/links")
	links.Use(middleware.Auth())
	links.GET("/", controller.IndexLink)
	links.GET("/new", controller.NewLink)
	links.POST("/", controller.CreateLink)
	links.GET("/:id/edit", controller.EditLink)
	links.POST("/:id", controller.UpdateLink)
	links.GET("/:id/delete", controller.DeleteLink)

	users = e.Group("/users/:loginId")
	// Links.
	folders := users.Group("/folders")
	folders.Use(middleware.Auth())
	folders.GET("/", controller.IndexFolder)
	folders.GET("/new", controller.NewFolder)
	folders.POST("/", controller.CreateFolder)
	folders.GET("/:id/edit", controller.EditFolder)
	folders.POST("/:id", controller.UpdateFolder)
	folders.GET("/:id/delete", controller.DeleteFolder)

	// Run server.
	e.Run(":8080")
}
