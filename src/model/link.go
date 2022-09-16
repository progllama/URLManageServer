package model

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	UserID   uint
	User     User
	FolderID uint

	Title string
	URL   string
}
