package models

import "github.com/jinzhu/gorm"

type LinkList struct {
	gorm.Model
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}
