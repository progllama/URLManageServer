package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MemSession struct {
	s sessions.Session
}

func (s *MemSession) HasUserId() bool {
	userId := s.s.Get(USER_ID)
	return userId != nil
}

func (s *MemSession) SetUserId(userId int) {
	s.s.Set(USER_ID, userId)
	err := s.s.Save()
	if err != nil {
		panic(err)
	}
}

func (s *MemSession) GetUserId() int {
	userId := s.s.Get(USER_ID)
	if userId == nil {
		panic(ErrKeyNotFound)
	}
	return userId.(int)
}

func (rs *MemSession) Clear() {
	rs.s.Clear()
	rs.s.Save()
}

func NewMemSession(ctx *gin.Context) *MemSession {
	session := sessions.Default(ctx)
	return &MemSession{
		s: session,
	}
}

type MemSessionFactory struct {
}

func (factory *MemSessionFactory) Create(c *gin.Context) Session {
	return NewMemSession(c)
}

func NewMemSessionFactory() *MemSessionFactory {
	return &MemSessionFactory{}
}
