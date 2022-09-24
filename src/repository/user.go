package repository

import (
	"errors"
	"url_manager/model"
)

var users []model.User = make([]model.User, 0)

func CreateUser(loginId string, password string) error {
	users = append(users, model.User{LoginId: loginId, Password: password})
	return nil
}

func GetUser(loginId string, password string) (*model.User, error) {
	for _, u := range users {
		if u.LoginId == loginId && u.Password == password {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func FindUser(loginId string) (*model.User, error) {
	for _, u := range users {
		if u.LoginId == loginId {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}
