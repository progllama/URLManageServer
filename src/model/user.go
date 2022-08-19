package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `json:"userId"`
	Password string `json:"password"`
	Links    []Link
}

func (u *User) Authenticate(cred *User) bool {
	if u.UserID != cred.UserID {
		return false
	}
	if u.Password != cred.Password {
		return false
	}
	return true
}

func NewUser(id int) *User {
	u := User{}
	u.ID = uint(id)
	return &u
}
