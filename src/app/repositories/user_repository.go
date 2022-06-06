package repositories

import (
	"url_manager/app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	All() ([]models.User, error)
	Find(int) (models.User, error)
	Add(models.User) error
	Update(models.User) error
	Remove(int) error
	Exists(int) (bool, error)
}

type userRepositoryImplPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImplPostgres{
		db,
	}
}

func (repo *userRepositoryImplPostgres) All() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *userRepositoryImplPostgres) Find(id int) (models.User, error) {
	var user models.User
	err := repo.db.Where("id=?", id).First(&user).Error
	return user, err
}

func (repo *userRepositoryImplPostgres) Add(user models.User) error {
	result := repo.db.Create(&user)
	return result.Error
}

func (repo *userRepositoryImplPostgres) Update(user models.User) error {
	result := repo.db.Save(&user)
	return result.Error
}

func (repo *userRepositoryImplPostgres) Remove(id int) error {
	result := repo.db.Delete(models.User{}, id)
	return result.Error
}

func (repo *userRepositoryImplPostgres) Exists(id int) (bool, error) {
	var user models.User
	result := repo.db.Where("id=?", id).Limit(1).Find(user)
	return result.RowsAffected > 0, result.Error
}
