package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"Name" json:"name" binding:"required" gorm:"unique;not null"`
	Email    string `form:"Email" json:"email" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
}
