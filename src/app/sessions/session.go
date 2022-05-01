package sessions

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

type SessionFactory interface {
	Create(c *gin.Context) Session
}

type RedisSessionFactory struct {
}

func (factory *RedisSessionFactory) Create(c *gin.Context) Session {
	return NewRedisSession(c)
}

type RedisSession struct {
	s sessions.Session
}

func (rs *RedisSession) HasUserId() bool {
	userId := rs.s.Get(rs.getUserIdKeyName())
	return userId != nil
}

func (rs *RedisSession) SetUserId(userId int) {
	rs.s.Set(rs.getUserIdKeyName(), userId)
	err := rs.s.Save()
	if err != nil {
		panic(err)
	}
}

func (rs *RedisSession) GetUserId() int {
	userId := rs.s.Get(rs.getUserIdKeyName())
	if userId == nil {
		panic(ErrKeyNotFound)
	}
	return userId.(int)
}

func (rs *RedisSession) Clear() {
	rs.s.Clear()
	rs.s.Save()
}

func (rs *RedisSession) getUserIdKeyName() string {
	// こういう書き方をすると値を書き換えるたびに再コンパイルが必要になってしまうが。このぐらいは問題なしとする。
	return "user_id"
}

func NewRedisSession(ctx *gin.Context) *RedisSession {
	session := sessions.Default(ctx)
	return &RedisSession{
		s: session,
	}
}
