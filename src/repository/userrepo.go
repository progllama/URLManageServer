package repository

import (
	"errors"
	"url_manager/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	All() ([]model.User, error)
	Get(int) (model.User, error)
	Create(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	Delete(int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) All() ([]model.User, error) {
	var users []model.User
	resultSet := repo.db.Select("id, name").Find(&users)
	return users, resultSet.Error
}

func (repo *userRepository) Get(id int) (model.User, error) {
	var user model.User
	resultSet := repo.db.Select("id, name").Where("id=?", id).First(&user)
	if errors.Is(resultSet.Error, gorm.ErrRecordNotFound) {
		return model.User{}, ErrItemNotFound
	}
	return user, resultSet.Error
}

func (repo *userRepository) Create(u model.User) (model.User, error) {
	resultSet := repo.db.Create(&u)
	if resultSet.Error != nil {
		return u, resultSet.Error
	}

	var created model.User
	resultSet = repo.db.Select("id, name").Where("name=?", u.Name).First(&created)
	return created, resultSet.Error
}

func (repo *userRepository) Update(u model.User) (model.User, error) {
	resultSet := repo.db.Where("id=?", u.ID).Save(&u)
	if resultSet.Error != nil {
		return u, resultSet.Error
	}
	var updated model.User
	resultSet = repo.db.Select("id, name").Where("id=?", u.ID).First(&updated)
	return updated, resultSet.Error
}

func (repo *userRepository) Delete(id int) error {
	u := model.User{}
	u.ID = uint(id)
	return repo.db.Delete(&u).Error
}
