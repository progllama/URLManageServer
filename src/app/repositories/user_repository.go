package repositories

import (
	"fmt"
	"url_manager/app/models"
	"url_manager/db"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(uint) (models.User, error)
	GetByName(string) (models.User, error)
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delte() error
}

type UserRepository struct {
}

func (urepo UserRepository) GetAll() ([]models.User, error) {
	db := db.GetDB()
	var user []models.User
	if err := db.Table("users").Select("name, id").Scan(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (urepo UserRepository) GetByID(id int) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	db.Table("users").Where("id = ?", id).First(&user)

	return user, nil
}

func (urepo UserRepository) GetByName(name string) (models.User, error) {
	db := db.GetDB()
	var user models.User

	if err := db.Table("users").Where("name=?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (urepo UserRepository) Create(c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var user models.User

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

func (_ UserRepository) Update(UserRepository) (models.User, error) {
	// db := db.GetDB()
	// var user models.User
	// if err := db.Where("id = ?", id).First(&user).Error; err != nil {
	// 	return user, err
	// }
	// if err := c.BindJSON(&user); err != nil {
	// 	return user, err
	// }
	// user.ID = uint(id)
	// db.Save(&user)

	// return user, nil
}

func (_ UserRepository) Delete(id int) error {
	db := db.GetDB()
	var user models.User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
