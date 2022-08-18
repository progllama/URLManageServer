package repository

import "url_manager/model"

type UserRepository struct {
}

func (repo *UserRepository) All(userID string) []model.User {
	panic("")
}

func (repo *UserRepository) Get(id int) model.User {
	panic("")
}

func (repo *UserRepository) Create(c model.User) model.User {
	panic("")
}

func (repo *UserRepository) Update(id int, c model.User) model.User {
	panic("")
}

func (repo *UserRepository) Delete(id int) {
	panic("")
}
