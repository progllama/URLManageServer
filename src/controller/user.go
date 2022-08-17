package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
