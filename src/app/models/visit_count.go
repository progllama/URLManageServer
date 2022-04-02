package models

import "github.com/jinzhu/gorm"

type VisitCount struct {
	gorm.Model
	URLID  int
	UserID int
	Count  int
}
