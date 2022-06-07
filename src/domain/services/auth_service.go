package services

import (
	"url_manager/app/models"
)

type AuthService interface {
	Login(models.Credential) LoginResponse
}

type LoginResponse struct {
	Code    int
	Body    interface{}
	Success bool
	UserId  int
}
