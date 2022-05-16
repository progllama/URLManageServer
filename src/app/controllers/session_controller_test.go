package controllers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_manager/app/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestNewSessionSuccess(t *testing.T) {
	// config.
	method := "GET"
	url := "/session/new"

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	//  Set dummy template.
	r.SetHTMLTemplate(template.Must(template.New("login.html").Parse("")))

	// testing preparation.
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Error(err)
	}

	// set route.
	ctrl := NewSessionController(nil)
	r.GET(url, ctrl.NewSession)

	// call
	r.ServeHTTP(w, req)

	// check status code.
	assert.Equal(t, 200, w.Code)
}

func TestCreateSessionFaultFormBind(t *testing.T) {
	// config.
	method := "POST"
	url := "/session"

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// testing preparation.
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Error(err)
	}

	// set route.
	ctrl := NewSessionController(nil)
	r.POST(url, ctrl.CreateSession)

	// call
	r.ServeHTTP(res, req)

	// check status code.
	assert.Equal(t, 500, res.Code)
}

func TestCreateSessionUserNoExist(t *testing.T) {
	// config.
	method := "POST"
	url := "/session"
	redirectUrl := "/users/new"
	body := bytes.NewBufferString("login_id=dummy&password=dummy")

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// testing preparation.
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Error(err)
	}

	// set route.
	repo := repositories.NewUserRepositoryMock()
	ctrl := NewSessionController(repo)
	r.POST(url, ctrl.CreateSession)

	// call
	r.ServeHTTP(res, req)

	// check status code.
	assert.Equal(t, 302, res.Code)
	assert.Equal(t, redirectUrl, res.Header().Get("Location"))
}

func TestCreateSessionAuthFaultWithCorrectNameAndWrongPass(t *testing.T) {
	// config.
	method := "POST"
	url := "/session"
	redirectUrl := "/users/new"
	body := bytes.NewBufferString("login_id=dummy&password=OK")

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// testing preparation.
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Error(err)
	}
	req.Header = map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	// set route.
	repo := repositories.NewUserRepositoryMock()
	repo.Create("test-name", "test-login-id", "OK")

	ctrl := NewSessionController(repo)
	r.POST(url, ctrl.CreateSession)

	// call
	r.ServeHTTP(res, req)

	// check status code.
	assert.Equal(t, 302, res.Code)
	assert.Equal(t, redirectUrl, res.Header().Get("Location"))
}

func TestCreateSessionAuthFaultWithWrongNameAndCorrectPass(t *testing.T) {
	// config.
	method := "POST"
	url := "/session"
	redirectUrl := "/users/new"
	body := bytes.NewBufferString("login_id=dummy&password=OK")

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	r := gin.New()

	// testing preparation.
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Error(err)
	}
	req.Header = map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	// set route.
	repo := repositories.NewUserRepositoryMock()
	repo.Create("test-name", "dummy", "NG")

	ctrl := NewSessionController(repo)
	r.POST(url, ctrl.CreateSession)

	// call
	r.ServeHTTP(res, req)

	// check status code.
	assert.Equal(t, 302, res.Code)
	assert.Equal(t, redirectUrl, res.Header().Get("Location"))
}

func TestCreateSessionSuccess(t *testing.T) {
	// config.
	method := "POST"
	url := "/session"
	redirectUrl := "/users/0"
	body := bytes.NewBufferString("login_id=dummy&password=OK")

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// make and register session store.
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// testing preparation.
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Error(err)
	}
	req.Header = map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	// set route.
	repo := repositories.NewUserRepositoryMock()
	repo.Create("test-name", "dummy", "OK")

	ctrl := NewSessionController(repo)
	r.POST(url, ctrl.CreateSession)

	// call
	r.ServeHTTP(res, req)

	// check status code.
	assert.Equal(t, 302, res.Code)
	assert.Equal(t, redirectUrl, res.Header().Get("Location"))
}

// セッションの履歴を管理できないの本当に消えてるかのテストができないない。
func TestDeleteSession(t *testing.T) {
	// config.
	method := "DELETE"
	url := "/session"
	body := bytes.NewBufferString("login_id=dummy&password=OK")

	// inactivate gin's log.
	gin.DefaultWriter = ioutil.Discard

	// Create new router
	r := gin.New()

	// make and register session store.
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// testing preparation.
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Error(err)
	}

	// set route.
	repo := repositories.NewUserRepositoryMock()

	ctrl := NewSessionController(repo)
	r.DELETE(url, ctrl.DestroySession)

	// call
	r.ServeHTTP(res, req)

	// check status code.
	assert.Equal(t, 200, res.Code)
}
