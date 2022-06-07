package services

import "github.com/gin-gonic/gin"

type AuthService interface {
	Login(ctx *gin.Context) ServiceResponse
	Logout(ctx *gin.Context) ServiceResponse
}

type ServiceResponse struct {
	Code int
	Body interface{}
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (service *authService) Login(ctx *gin.Context) ServiceResponse {
	panic("")
}

func (service *authService) Logout(ctx *gin.Context) ServiceResponse {
	panic("")
}
