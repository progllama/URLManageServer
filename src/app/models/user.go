package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" gorm:"size:100"`
}
