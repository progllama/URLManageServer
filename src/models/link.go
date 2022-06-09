package models

import "github.com/jinzhu/gorm"

type Link struct {
	gorm.Model
	UserID int    `json:"user_id"`
	Url    string `json:"url" gorm:"not null"`
	Title  string `json:"title" gorm:"not null"`
}
