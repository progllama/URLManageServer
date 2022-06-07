package models

import "github.com/jinzhu/gorm"

type Url struct {
	gorm.Model
	OwnerId     int
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string ``
	Note        string ``
	VisitCount  int
	User        User `gorm:"foreignKey:OwnerId"`
}

func NewUrl() *Url {
	return &Url{}
}
