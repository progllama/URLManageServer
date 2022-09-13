package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LogoutService interface {
	Logout(id string) error
}

type LogoutServiceFactory interface {
	NewLogoutService(*gin.Context) LogoutService
}

func NewLogoutService(context *gin.Context) LogoutService {
	return &logoutService{
		context,
	}
}

type logoutService struct {
	context *gin.Context
}

func (service *logoutService) Logout(loginId string) (err error) {
	session := sessions.Default(service.context)
	session.Clear()
	err = session.Save()
	return
}
