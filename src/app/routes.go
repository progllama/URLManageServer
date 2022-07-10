package app

import "github.com/gin-gonic/gin"

func Routes(engine *gin.Engine) {
	engine.LoadHTMLGlob("template/*")
	engine.GET("/", Index("index.html"))
	engine.GET("/urls", Index("index.html"))
	engine.GET("/urls/new", New("new.html"))
	engine.POST("/urls", Create())
}
