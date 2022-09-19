package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetLink(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.New()
	router.GET("/", GetLink)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "", res.Body.String())
}

func TestCreateLink(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.New()
	router.POST("/", CreateLink)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "", res.Body.String())
}
