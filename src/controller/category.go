package controllers

import (
	"net/http"
	"url_manager/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	var categories []model.Category
	ctx.JSON(http.StatusOK, categories)
}

func GetCategory(ctx *gin.Context) {
	var category model.Category
	ctx.JSON(http.StatusOK, category)
}

func CreateCategory(ctx *gin.Context) {
	var category model.Category
	ctx.JSON(http.StatusOK, category)
}

func UpdateCategory(ctx *gin.Context) {
	var category model.Category
	ctx.JSON(http.StatusOK, category)
}

func DeleteCategory(ctx *gin.Context) {
	empty := gin.H{}
	ctx.JSON(http.StatusOK, empty)
}
