package models

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	URL         string   ``
	Title       string   ``
	Tag         []Tag    ``
	Description string   ``
	Memos       []Memo   ``
	Status      []Status ``
	VisitCount  uint     ``
	Icon        string   ``
	URLGroupID  uint     ``
}

type Tag struct {
	gorm.Model
	URLID uint
	Name  string
}
