package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RedisSession struct {
	s sessions.Session
}

func (s *RedisSession) ID() string {
	return s.s.ID()
}

func (rs *RedisSession) HasUserId() bool {
	userId := rs.s.Get(USER_ID)
	return userId != nil
}

func (rs *RedisSession) SetUserId(userId int) {
	rs.s.Set(USER_ID, userId)
	err := rs.s.Save()
	if err != nil {
		panic(err)
	}
}

func (rs *RedisSession) GetUserId() int {
	userId := rs.s.Get(USER_ID)
	if userId == nil {
		panic(ErrKeyNotFound)
	}
	return userId.(int)
}

func (rs *RedisSession) Clear() {
	rs.s.Clear()
	rs.s.Save()
}

func NewRedisSession(ctx *gin.Context) *RedisSession {
	session := sessions.Default(ctx)
	return &RedisSession{
		s: session,
	}
}

type RedisSessionFactory struct {
}

func (factory *RedisSessionFactory) Create(c *gin.Context) Session {
	return NewRedisSession(c)
}

func NewRedisSessionFactory() *RedisSessionFactory {
	return &RedisSessionFactory{}
}
