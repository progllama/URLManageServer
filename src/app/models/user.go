package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" json:"password" binding:"required" gorm:"size:100"`
}
