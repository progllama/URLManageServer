package models

import "github.com/jinzhu/gorm"

type Url struct {
	gorm.Model
	OwnerId     int
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string ``
	Note        string ``
	VisitCount  int
}

func NewUrl() *Url {
	return &Url{}
}

type Comment struct {
	gorm.Model
	Owner User
	Url   Url
	Text  string
}

type CommentRelationShip struct {
	gorm.Model
	Parent Comment
	Child  Comment
}

type MasterTag struct {
	gorm.Model
	Name string
}

type TagMap struct {
	gorm.Model
	Url       Url
	MasterTag MasterTag
}
