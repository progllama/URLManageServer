package model

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	URLID   uint
	Order   int
	Content string
}
