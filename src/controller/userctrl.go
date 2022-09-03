package controllers

import (
	"net/http"
	"strconv"
	"url_manager/model"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	repo := getUserRepo()
	var create model.User
	ctx.BindJSON(&create)
	created := repo.Create(create)
	ctx.JSON(http.StatusOK, created)
}

func UpdateUser(ctx *gin.Context) {
	repo := getUserRepo()
	var update model.User
	ctx.BindJSON(&update)
	userId, _ := strconv.Atoi(ctx.Param(":id"))
	updated := repo.Update(userId, update)
	ctx.JSON(http.StatusOK, updated)
}

func DeleteUser(ctx *gin.Context) {
	repo := getUserRepo()
	id, _ := strconv.Atoi(ctx.Param(":id"))
	repo.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{})
}

func getUserRepo() *repository.UserRepository {
	return &repository.UserRepository{}
}
