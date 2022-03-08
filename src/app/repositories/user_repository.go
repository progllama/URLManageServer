package repositories

import (
	"url_manager/app/models"
	"url_manager/db"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

type User models.User

// Userをそのまま返しているのでテーブル構造がバレてしまう。

// GetAll is get all User
func (_ UserRepository) GetAll() ([]User, error) {
	db := db.GetDB()
	var user []User
	if err := db.Table("users").Select("name, id").Scan(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateModel is create User model
func (_ UserRepository) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var user User

	// c.Request.ParseForm()
	// user.Name = c.Request.FormValue("Name")
	// user.Email = c.Request.FormValue("Email")
	// user.Password = c.Request.FormValue("Password")
	c.Bind(&user)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12) // 2 ^ 12 回　ストレッチ回数

	if err != nil {
		user.Password = string(hashedPass)
	}

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByID is get a User by ID
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

// UpdateByID is update a User
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

// DeleteByID is delete a User by ID
func (_ UserRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var user User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
