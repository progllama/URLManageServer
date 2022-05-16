package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_manager/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestUserApiIndexSuccess(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	repo := repositories.NewUserRepositoryMock()
	api := NewUserApi(repo)
	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestUserApiIndexSuccessAny(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	repo := repositories.NewUserRepositoryMock()
	repo.Create("test-name", "login-id", "password")
	api := NewUserApi(repo)
	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	want := `["name":"test-name"}]`
	assert.Equal(t, want, w.Body.String())
}

func TestUserApiIndexFail(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	repo := repositories.NewUserRepositoryMock()
	repo.Error = errors.New("Dummy error.")
	api := NewUserApi(repo)
	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func NewTestRouter() *gin.Engine {
	gin.DefaultWriter = ioutil.Discard
	return gin.New()
}
