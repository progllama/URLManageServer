package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required" gorm:"unique;not null"`
	LoginId  string `form:"login_id" json:"login_id" binding:""`
	Password string `form:"password" json:"password" binding:"required" gorm:"size:100"`
	Urls     []Url  `gorm:"foreignKey:OwnerId"`
}

// TODO パスワードのハッシュ化はこの構造体の責務"オブジェクトの永続化"でないので移動する。
func (user *User) GenerateHashFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12) // 2 ^ 12 回　ストレッチ回数
	return string(hash), err
}
