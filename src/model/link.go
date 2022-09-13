package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	UserID   uint
	User     User
	FolderID uint
	Title    string `json:"title"`
	URL      string `json:"url"`
}

func (link *Link) IsValid() bool {
	validator := validator.New()
	err := validator.Struct(link)
	return err != nil
}
