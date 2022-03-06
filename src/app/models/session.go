package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ISession interface {
	ID() (string, error)
	Get(string) (string, error)
	Set(string, string) (string, error)
	Delte(string)
	Clear() error
	Save() error
}

type Session struct {
	session sessions.Session
}

func NewSession(ctx *gin.Context) Session {
	return Session{
		sessions.Default(ctx),
	}
}

func (s Session) ID() (string, error) {
	return s.session.ID(), nil
}

func (s Session) Get(k string) (string, error) {
	return s.Get(k)
}

func (s Session) Set(k string, v string) (string, error) {
	return s.Set(k, v)
}

func (s Session) Delete(k string) (string, error) {
	return s.Delete(k)
}

func (s Session) Clear() (string, error) {
	return s.Clear()
}

func (s Session) Save() error {
	return s.Save()
}
