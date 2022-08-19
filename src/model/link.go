package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	UserID uint
	User   User
	Title  string `json:"title"`
	URL    string `json:"url"`
}

func NewLink(id int) *Link {
	l := Link{}
	l.ID = uint(id)
	return &l
}
