package route

import (
	"net/http"
	controllers "url_manager/controller"
	"url_manager/session"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	RegisterRoot(e)
	RegisterSignInAndOut(e)
	RegisterUserRoutes(e)
	RegisterLinkRoutes(e)
}

func RegisterRoot(e *gin.Engine) {
	e.LoadHTMLGlob("public/index.html")
	e.Static("/static", "./public/static")
	e.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
}

func RegisterSignInAndOut(e *gin.Engine) {
	e.POST("/sign_in", session.Middleware(), controllers.Login)
	e.DELETE("/sign_out", session.Middleware(), controllers.Logout)
}

func RegisterUserRoutes(e *gin.Engine) {
	g := e.Group("/users")
	g.Use(session.Middleware())
	g.POST("/", controllers.CreateUser)
	g.PUT("/:user_id", controllers.Authenticate, controllers.UpdateUser)
	g.DELETE("/:user_id", controllers.Authenticate, controllers.DeleteUser)
}

func RegisterLinkRoutes(e *gin.Engine) {
	g := e.Group("users/:user_id/links")
	g.Use(session.Middleware())
	g.Use(controllers.Authenticate)
	g.GET("/", controllers.GetLinks)
	g.POST("/", controllers.CreateLink)
	g.PUT("/:id", controllers.UpdateLink)
	g.DELETE("/:id", controllers.DeleteLink)
}
