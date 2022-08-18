package controllers

import (
	"net/http"
	"strconv"
	"url_manager/model"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	repo := getUserRepo()
	userId := ctx.Param(":user_id")
	users := repo.All(userId)
	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	repo := getUserRepo()
	id, _ := strconv.Atoi(ctx.Param(":id"))
	user := repo.Get(id)
	ctx.JSON(http.StatusOK, user)
}

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
	empty := gin.H{}
	ctx.JSON(http.StatusOK, empty)
}

func getUserRepo() *repository.UserRepository {
	return &repository.UserRepository{}
}
