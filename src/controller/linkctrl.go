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
	userId, _ := strconv.Atoi(ctx.Param("user_id"))
	links := repo.All(userId)
	ctx.JSON(http.StatusOK, links)
}

func CreateLink(ctx *gin.Context) {
	repo := getLinkRepo()
	userId, _ := strconv.Atoi(ctx.Param("user_id"))
	var create model.Link
	ctx.BindJSON(&create)
	created := repo.Create(userId, create)
	ctx.JSON(http.StatusOK, created)
}

func UpdateLink(ctx *gin.Context) {
	repo := getLinkRepo()
	id, _ := strconv.Atoi(ctx.Param("id"))
	var update model.Link
	ctx.BindJSON(&update)
	updated := repo.Update(id, update)
	ctx.JSON(http.StatusOK, updated)
}

func DeleteLink(ctx *gin.Context) {
	repo := getLinkRepo()
	id, _ := strconv.Atoi(ctx.Param("id"))
	repo.Delete(id)
	ctx.JSON(http.StatusOK, emptyBody)
}

func getLinkRepo() *repository.LinkRepository {
	return &repository.LinkRepository{}
}
