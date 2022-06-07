package api

import (
	"url_manager/interface/web/gin/services"

	"github.com/gin-gonic/gin"
)

type AuthApi interface {
	Login(*gin.Context)
	Logout(*gin.Context)
}

func NewAuthApi(s services.AuthService) AuthApi {
	return &authApi{}
}

type authApi struct {
	service services.AuthService
}

func (api *authApi) Login(ctx *gin.Context) {
	response := api.service.Login(ctx)
	ctx.JSON(response.Code, response.Body)
}

func (api *authApi) Logout(ctx *gin.Context) {
	response := api.service.Logout(ctx)
	ctx.JSON(response.Code, response.Body)
}
