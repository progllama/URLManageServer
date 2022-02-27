package models

import "github.com/jinzhu/gorm"

type URLGroup struct {
	gorm.Model
	UserID uint
	ID     uint
	Title  string
}
