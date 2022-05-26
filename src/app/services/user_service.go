package services

import "url_manager/app/models"

type UserService interface {
	FindUsers() UserServiceResponse
	FindUser(string) UserServiceResponse
	Create(models.User) UserServiceResponse
	Update(models.User) UserServiceResponse
	Delete(models.User) UserServiceResponse
}

type UserServiceResponse interface {
	Code() int
	Body() interface{}
}
