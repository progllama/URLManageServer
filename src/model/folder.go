package model

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name      string `json:"name" validate:"required,min=1,max=256,alphanum" gorm:"size:256;not null;check:(LENGTH(name) <= 256 AND LENGTH(name) > 0)"`
	ManagerID *uint
	Folders   []Folder `gorm:"foreignkey:ManagerID"`
	Links     []Link
	UserID    uint
}
