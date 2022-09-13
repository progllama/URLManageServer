package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginServiceFactory interface {
	NewLoginService(*gin.Context) LoginService
}

type LoginService interface {
	IsLogin(loginId string) bool
	Login(loginId string) error
}

func NewLoginServiceFactory() LoginServiceFactory {
	return &loginServiceFactory{}
}

func NewLoginService(context *gin.Context) LoginService {
	return &loginService{
		context,
	}
}

type loginService struct {
	context *gin.Context
}

func (service *loginService) Login(loginId string) (err error) {
	session := sessions.Default(service.context)
	session.Set(loginIdKey, loginId)
	err = session.Save()
	return
}

func (service *loginService) IsLogin(loginId string) bool {
	s := sessions.Default(service.context)
	sessionLoginId := s.Get(loginIdKey)

	if sessionLoginId == nil {
		return false
	}

	if sessionLoginId.(string) != loginId {
		return false
	}

	return true
}

type loginServiceFactory struct {
}

func (factory *loginServiceFactory) NewLoginService(context *gin.Context) LoginService {
	return NewLoginService(context)
}
