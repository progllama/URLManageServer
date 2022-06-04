package models

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" binding:"required" gorm:"unique;not null"`
	LoginId   string `json:"login_id" binding:""`
	Password  string `json:"password" binding:"required" gorm:"size:100"`
	Urls      []Url  `gorm:"foreignKey:OwnerId"`
}

func (user *User) Authenticate(loginId string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// TODO パスワードのハッシュ化はこの構造体の責務"オブジェクトの永続化"でないので移動する。
// 本当にそう？
func (user *User) GenerateHashFromPassword(password string) (string, error) {
	strings.Split("", " ")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12) // 2 ^ 12 回　ストレッチ回数
	return string(hash), err
}
