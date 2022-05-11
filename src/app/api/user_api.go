package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

func (api *UserApi) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
