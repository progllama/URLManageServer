package models

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	URL    string `json:"url"`
	Title  string `json:"title"`
	UserID int    `json:"userId"`
	// Tag         []Tag    ``
	// Description string   ``
	// Memos       []Memo   ``
	// Status      []Status ``
	// VisitCount  uint     ``
	// Icon        string   ``
	// URLGroupID  uint     ``
}

// type Tag struct {
// 	gorm.Model
// 	URLID uint
// 	Name  string
// }
