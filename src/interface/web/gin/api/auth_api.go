package api

import (
	"net/http"
	"url_manager/interface/web/gin/services"

	"github.com/gin-contrib/sessions"
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
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"logout": "success"})
}

func extractCredential(ctx *gin.Context) {

}
