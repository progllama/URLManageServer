package repositories

import (
	"fmt"
	"url_manager/app/models"
	"url_manager/db"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	GetByID(uint) (models.User, error)
	GetByName(string) (models.User, error)
	GetAll() (models.User, error)
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delte() (models.User, error)
}

type UserRepository struct{}

type User models.User

func (_ UserRepository) GetAll() ([]User, error) {
	db := db.GetDB()
	var user []User
	if err := db.Table("users").Select("name, id").Scan(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (_ UserRepository) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var user User

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12) // 2 ^ 12 回　ストレッチ回数

	if err != nil {
		user.Password = string(hashedPass)
	}

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (_ UserRepository) GetByID(id int) (User, error) {
	db := db.GetDB()
	var user User
	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	db.Table("users").Where("id = ?", id).First(&user)

	return user, nil
}

func (_ UserRepository) GetByName(name string) (User, error) {
	db := db.GetDB()
	var user User

	if err := db.Table("users").Where("name=?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (_ UserRepository) UpdateByID(id int, c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	user.ID = uint(id)
	db.Save(&user)

	return user, nil
}

func (_ UserRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var user User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
