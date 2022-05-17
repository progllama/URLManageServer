package middlewares

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_manager/app/session"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestLoggdInUser(t *testing.T) {
	// config.
	method := "GET"
	url := "/"

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// testing preparation.
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Error(err)
	}

	f := &SessionFactoryMock{}
	f.dummy = &SessionMock{}
	f.dummy.hasUseId = true
	middleware := NewLoginRequireMiddleware(f, RedirectToLoginPage, DoNothing)
	r.Use(middleware.RequireLogin())
	r.GET(url, func(c *gin.Context) {
		c.String(http.StatusOK, "💙💚💛")
	})

	// call
	r.ServeHTTP(w, req)

	// check status code.
	assert.Equal(t, 302, w.Code)
}

func TestNotLoggdInUser(t *testing.T) {
	// config.
	method := "GET"
	url := "/"

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// testing preparation.
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Error(err)
	}

	f := &SessionFactoryMock{}
	f.dummy = &SessionMock{}
	f.dummy.hasUseId = false
	middleware := NewLoginRequireMiddleware(f, RedirectToLoginPage, DoNothing)
	r.Use(middleware.RequireLogin())
	r.GET(url, func(c *gin.Context) {
		c.String(http.StatusOK, "💙💚💛")
	})

	// call
	r.ServeHTTP(w, req)

	// check status code.
	assert.Equal(t, 200, w.Code)
}

type SessionFactoryMock struct {
	dummy *SessionMock
}

func (f *SessionFactoryMock) Create(_ *gin.Context) session.Session {
	return f.dummy
}

type SessionMock struct {
	session.Session
	hasUseId bool
}

func (s *SessionMock) HasUserId() bool {
	return s.hasUseId
}
