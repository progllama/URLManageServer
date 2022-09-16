package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	LoginId string
	IndexId int
}
