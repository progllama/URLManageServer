package models

import "gorm.io/gorm"

type LinkListRelation struct {
	gorm.Model
	UserID       int `json:"user_id"`
	ParentListID int `json:"parent_list_id"`
	ChildListID  int `json:"child_list_id"`
}
