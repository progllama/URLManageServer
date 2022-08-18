package controllers

import (
	"net/http"
	"strconv"
	"url_manager/model"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func GetLinks(ctx *gin.Context) {
	repo := getLinkRepo()
	userId := ctx.Param(":user_id")
	links := repo.All(userId)
	ctx.JSON(http.StatusOK, links)
}

func GetLink(ctx *gin.Context) {
	repo := getLinkRepo()
	id, _ := strconv.Atoi(ctx.Param(":id"))
	link := repo.Get(id)
	ctx.JSON(http.StatusOK, link)
}

func CreateLink(ctx *gin.Context) {
	repo := getLinkRepo()
	var create model.Link
	ctx.BindJSON(&create)
	created := repo.Create(create)
	ctx.JSON(http.StatusOK, created)
}

func UpdateLink(ctx *gin.Context) {
	repo := getLinkRepo()
	var update model.Link
	ctx.BindJSON(&update)
	updated := repo.Create(update)
	ctx.JSON(http.StatusOK, updated)
}

func DeleteLink(ctx *gin.Context) {
	repo := getLinkRepo()
	id, _ := strconv.Atoi(ctx.Param(":id"))
	repo.Delete(id)
	empty := gin.H{}
	ctx.JSON(http.StatusOK, empty)
}

func getLinkRepo() *repository.LinkRepository {
	return &repository.LinkRepository{}
}
