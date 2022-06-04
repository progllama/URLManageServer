package repositories

import (
	"url_manager/domain/models"

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

type UserRepositoryImplPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImplPostgres{
		db,
	}
}

func (repo *UserRepositoryImplPostgres) All() ([]models.User, error) {
	var users []models.User
	repo.db.Select("?, ?").Find(&users)
	return users
}

func (repo *UserRepositoryImplPostgres) Find(id int) (models.User, error) {
	panic("")
}

func (repo *UserRepositoryImplPostgres) Add(user models.User) error {
	panic("")
}

func (repo *UserRepositoryImplPostgres) Update(user models.User) error {
	panic("")
}

func (repo *UserRepositoryImplPostgres) Remove(id int) error {
	panic("")
}

func (repo *UserRepositoryImplPostgres) Exists(id int) (bool, error) {
	panic("")
}
