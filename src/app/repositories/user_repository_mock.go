package repositories

import (
	"errors"
	"time"
	"url_manager/app/models"
)

type UserRepositoryMock struct {
	users []models.User
	Error error
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		make([]models.User, 0),
		nil,
	}
}

func (r *UserRepositoryMock) All() ([]models.User, error) {
	return r.users, r.Error
}

func (r *UserRepositoryMock) AllIdName() ([]SafeUser, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepositoryMock) FindById(_ int) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepositoryMock) FindByName(_ string) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepositoryMock) FindByLoginId(loginId string) (models.User, error) {
	for _, u := range r.users {
		if u.LoginId == loginId {
			return u, nil
		}
	}
	return models.User{}, errors.New("User not Found")
}

func (r *UserRepositoryMock) Create(name string, loginId string, password string) error {
	u := models.User{
		ID:        len(r.users),
		CreatedAt: time.Date(2022, 5, 10, 22, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2022, 5, 10, 22, 0, 0, 0, time.Local),
		Name:      name,
		LoginId:   loginId,
		Password:  "",
	}
	u.Password, _ = u.GenerateHashFromPassword(password)
	r.users = append(r.users, u)
	return r.Error
}

func (r *UserRepositoryMock) Update(id int, name string, login string, password string) error {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepositoryMock) Delete(id int) error {
	panic("not implemented") // TODO: Implement
}

// HasUserId(int) (bool, error)
// HasUserLoginId(string) (bool, error)
func (r *UserRepositoryMock) HasUserName(_ string) (bool, error) {
	panic("not implemented") // TODO: Implement
}
