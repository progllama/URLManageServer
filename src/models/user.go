package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OpenID string `json:"open_id" gorm:"unique;not null"`
}
