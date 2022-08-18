package controllers

import (
	"net/http"
	"strconv"
	"url_manager/model"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	repo := getCategoryRepo()
	userId := ctx.Param(":user_id")
	categories := repo.All(userId)
	ctx.JSON(http.StatusOK, categories)
}

func GetCategory(ctx *gin.Context) {
	repo := getCategoryRepo()
	id, _ := strconv.Atoi(ctx.Param(":id"))
	category := repo.Get(id)
	ctx.JSON(http.StatusOK, category)
}

func CreateCategory(ctx *gin.Context) {
	repo := getCategoryRepo()
	var create model.Category
	ctx.BindJSON(&create)
	created := repo.Create(create)
	ctx.JSON(http.StatusOK, created)
}

func UpdateCategory(ctx *gin.Context) {
	repo := getCategoryRepo()
	var update model.Category
	ctx.BindJSON(&update)
	userId, _ := strconv.Atoi(ctx.Param(":id"))
	updated := repo.Update(userId, update)
	ctx.JSON(http.StatusOK, updated)
}

func DeleteCategory(ctx *gin.Context) {
	repo := getCategoryRepo()
	id, _ := strconv.Atoi(ctx.Param(":id"))
	repo.Delete(id)
	empty := gin.H{}
	ctx.JSON(http.StatusOK, empty)
}

func getCategoryRepo() *repository.CategoryRepository {
	return &repository.CategoryRepository{}
}
