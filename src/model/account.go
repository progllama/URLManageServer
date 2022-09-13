package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	OpenId string `json:"openId" validate:"required,max=256" gorm:"unique;size:256:not null;check:open_id <> ''"`
	UserID uint   `json:"-" validate:"required" gorm:"unique;not null;user_id <> ''"`
}
