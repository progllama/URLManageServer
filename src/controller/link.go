package controllers

import (
	"net/http"
	"url_manager/model"

	"github.com/gin-gonic/gin"
)

func GetLinks(ctx *gin.Context) {
	var links []model.Link
	ctx.JSON(http.StatusOK, links)
}

func GetLink(ctx *gin.Context) {
	var link model.Link
	ctx.JSON(http.StatusOK, link)
}

func CreateLink(ctx *gin.Context) {
	var link model.Link
	ctx.JSON(http.StatusOK, link)
}

func UpdateLink(ctx *gin.Context) {
	var link model.Link
	ctx.JSON(http.StatusOK, link)
}

func DeleteLink(ctx *gin.Context) {
	empty := gin.H{}
	ctx.JSON(http.StatusOK, empty)
}
