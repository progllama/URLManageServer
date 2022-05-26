package models

import "github.com/jinzhu/gorm"

type UrlDict struct {
	gorm.Model
	Owner User `gorm:""`
	Url   []Url
}
