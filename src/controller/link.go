package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLinks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func GetLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func CreateLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func UpdateLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func DeleteLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
