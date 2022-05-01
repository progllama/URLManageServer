package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	router := gin.Default()
	ctrl := NewUserController()

	router.GET("/users/new", ctrl.New)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/new", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
