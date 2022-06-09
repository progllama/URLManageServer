package models

import "github.com/jinzhu/gorm"

type LinkListRelation struct {
	gorm.Model
	UserID          int `json:"user_id"`
	ParentUrlListID int `json:"parent_url_list_id"`
	ChildUrlListID  int `json:"child_url_list_id"`
}
