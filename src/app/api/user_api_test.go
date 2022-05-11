package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestUserApiIndex(t *testing.T) {
	router := gin.New()
	route := "/api/users"

	api := NewUserApi()
	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{}", w.Body.String())
}
