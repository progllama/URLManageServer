package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Title string `json:"title"`
	URL   string `json:"url"`
}
