package services

import "url_manager/domain/models"

type UserService interface {
	FindUsers() UserServiceResponse
	FindUser(string) UserServiceResponse
	Create(models.User) UserServiceResponse
	Update(models.User) UserServiceResponse
	Delete(string) UserServiceResponse
}

type UserServiceResponse interface {
	Code() int
	Body() interface{}
}
