package session

import (
	"errors"
)

var (
	ErrKeyNotFound = errors.New("キーが見つかりませんでした。")
)

type Session interface {
	HasUserId() bool
	SetUserId(int)
	GetUserId() int
	Clear()
}
