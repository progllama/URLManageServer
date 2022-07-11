package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	UserID     int    //`json:"user_id"`
	LinkListID int    //`json:"link_list_id"`
	Url        string `json:"url" gorm:"not null"`
	Title      string `json:"title" gorm:"not null"`
}
