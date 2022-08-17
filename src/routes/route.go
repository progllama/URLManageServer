package route

import (
	"net/http"
	controllers "url_manager/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	RegisterRoot(e)
	RegisterUserRoutes(e)
	RegisterLinkRoutes(e)
}

func RegisterRoot(e *gin.Engine) {
	e.LoadHTMLGlob("public/index.html")
	e.Static("/static", "./public/static")
	e.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
}

func RegisterUserRoutes(e *gin.Engine) {
	g := e.Group("/users")
	g.GET("/", controllers.GetUsers)
	g.GET("/:id", controllers.GetUser)
	g.POST("/", controllers.CreateUser)
	g.PUT("/:id", controllers.UpdateUser)
	g.DELETE("/:id", controllers.DeleteUser)
}

func RegisterLinkRoutes(e *gin.Engine) {
	g := e.Group("/links")
	g.GET("/", controllers.GetLinks)
	g.GET("/:id", controllers.GetLink)
	g.POST("/", controllers.CreateLink)
	g.PUT("/:id", controllers.UpdateLink)
	g.DELETE("/:id", controllers.DeleteLink)
}
