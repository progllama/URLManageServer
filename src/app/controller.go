package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var links []Link

func Index(template string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, template, gin.H{"links": links})
	}
}

func New(template string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, template, gin.H{})
	}
}

func Create() func(*gin.Context) {
	return func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		url := ctx.PostForm("url")

		links = append(links, Link{title, url})

		ctx.Redirect(302, "/urls")
	}
}
