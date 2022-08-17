package controllers

import (
	"net/http"
	"url_manager/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var users []model.User
	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	var user model.User
	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var user model.User
	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	var user model.User
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	empty := gin.H{}
	ctx.JSON(http.StatusOK, empty)
}
