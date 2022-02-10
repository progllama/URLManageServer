package repository

import (
	"time"
	"url_manager/db"
	"url_manager/models"

	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

type User models.User

type UserProfile struct {
	Name string
	Id   int
}

// GetAll is get all User
func (_ UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var user_profiles []UserProfile
	if err := db.Table("users").Select("name, id").Scan(&user_profiles).Error; err != nil {
		return nil, err
	}
	return user_profiles, nil
}

// CreateModel is create User model
func (_ UserRepository) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var user User
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}

	user.CreatedAt = time.Now()
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByID is get a User by ID
func (_ UserRepository) GetByID(id int) (UserProfile, error) {
	db := db.GetDB()
	var user_profile UserProfile
	if err := db.Table("users").Where("id = ?", id).First(&user_profile).Error; err != nil {
		return user_profile, err
	}
	db.Where("id = ?", id).First(&user_profile)

	return user_profile, nil
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
