package session

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrKeyNotFound = errors.New("キーが見つかりませんでした。")
)

const USER_ID = "user_id"

type Session interface {
	ID() string
	HasUserId() bool
	SetUserId(int)
	GetUserId() int
	Clear()
}

type SessionFactory interface {
	Create(c *gin.Context) Session
}
