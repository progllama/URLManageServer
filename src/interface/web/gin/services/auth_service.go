package services

import (
	"net/http"
	"url_manager/domain/models"
	"url_manager/domain/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx *gin.Context) ServiceResponse
	Logout(ctx *gin.Context) ServiceResponse
}

type ServiceResponse struct {
	Code int
	Body interface{}
}

type authService struct {
	service services.AuthenticationService
}

func NewAuthService() AuthService {
	return &authService{}
}

func (service *authService) Login(ctx *gin.Context) ServiceResponse {
	var credential models.Credential
	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"message": err}}
	}

	ok, err := service.service.Authenticate(credential)
	if err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"message": err}}
	}

	if !ok {
		return ServiceResponse{http.StatusNotFound, gin.H{"message": "wrong credential"}}
	}

	// userId := service.FindUserByCredential(credential)
	userId := "abcd1234"

	session := sessions.Default(ctx)
	session.Set("user_id", userId)
	session.Save()

	return ServiceResponse{http.StatusOK, gin.H{"message": "success"}}
}

func (service *authService) Logout(ctx *gin.Context) ServiceResponse {
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"message": err}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "success"}}
}
