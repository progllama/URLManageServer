package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string   `json:"name" validate:"required,min=1,max=256" gorm:"unique;size:256:not null;check:name <> ''"`
	Account Account  `json:"-" validate:"-"`
	Index   Folder   `json:"-" validate:"-"`
	Folders []Folder `json:"-" validate:"-"`
	Links   []Link   `json:"-" validate:"-"`
}
