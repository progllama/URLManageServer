package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Entry(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
