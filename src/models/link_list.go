package models

import "gorm.io/gorm"

type LinkList struct {
	gorm.Model
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}
