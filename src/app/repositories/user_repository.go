package repositories

import (
	"errors"
	"fmt"
	"log"
	"url_manager/app/models"
	"url_manager/db"

	"gorm.io/gorm"
)

type IUserRepository interface {
	All() ([]models.User, error)
	AllIdName() ([]SafeUser, error)
	FindById(int) (models.User, error)
	FindByName(string) (models.User, error)
	FindByLoginId(string) (models.User, error)
	Create(name, loginId, password string) error
	Update(id int, name, login, password string) error
	Delete(id int) error
	// HasUserId(int) (bool, error)
	HasUserName(string) (bool, error)
	// HasUserLoginId(string) (bool, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (repo *UserRepository) All() ([]models.User, error) {
	var users []models.User
	dbResult := repo.db.Table("users").Select("*").Find(&users)

	if dbResult.Error == nil {
		return users, nil
	} else {
		return []models.User{}, dbResult.Error
	}
}

type SafeUser struct {
	Name string
	ID   string
}

func (repo *UserRepository) AllIdName() ([]SafeUser, error) {
	var users []SafeUser
	dbResult := repo.db.Table("users").Find(&users)

	if dbResult.Error != nil {
		return make([]SafeUser, 0), dbResult.Error
	}

	return users, nil
}

func (repo *UserRepository) FindById(id int) (models.User, error) {
	fmt.Println(id)
	var user models.User
	dbResult := repo.db.Where("id=?", id).First(&user)

	log.Println(id)

	if dbResult.Error == nil {
		return user, nil
	} else {
		return models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	dbResult := repo.db.Where("name=?", name).First(&user)

	if dbResult.Error == nil {
		return user, nil
	} else {
		return models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) FindByLoginId(loginId string) (models.User, error) {
	var user models.User
	dbResult := repo.db.Where("login_id=?", loginId).First(&user)

	if dbResult.Error == nil {
		return user, nil
	} else {
		return models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) Create(name string, loginId string, password string) error {
	user := models.User{
		Name:    name,
		LoginId: loginId,
	}
	hashPass, err := user.GenerateHashFromPassword(password)
	if err != nil {
		log.Println("Fail to generate hash.")
		return err
	}
	user.Password = hashPass

	dbResult := repo.db.Create(&user)
	return dbResult.Error
}

func (repo *UserRepository) Update(id int, name string, loginId string, password string) error {
	u := models.User{
		ID:       id,
		Name:     name,
		LoginId:  loginId,
		Password: password,
	}
	dbResult := repo.db.Model(&models.User{}).Updates(u)

	if dbResult != nil {
		return dbResult.Error
	} else {
		return nil
	}
}

func (repo *UserRepository) Delete(id int) error {
	var user models.User
	dbResult := repo.db.Where("id=?", id).Delete(&user)
	return dbResult.Error
}

func (repo *UserRepository) HasUserName(name string) (bool, error) {
	var users []models.User
	dbResult := repo.db.Select("id").Where("name=?", name).Limit(1).Take(&users)
	if !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) && dbResult.Error != nil {
		return false, dbResult.Error
	}

	if len(users) > 0 {
		log.Println("records > 0")
		return true, nil
	} else {
		log.Println("record <= 0")
		log.Println(len(users))
		return false, nil
	}
}
