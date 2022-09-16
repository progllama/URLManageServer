package model

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name string

	Parent int
	UserID uint
	Index  bool
}
