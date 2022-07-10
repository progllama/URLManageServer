package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	stopGinLogging()
	router := gin.New()

	router.LoadHTMLGlob("template/**")
	router.GET("/home", Index("dummy.html"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/home", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())
}

func TestIndexList(t *testing.T) {
	stopGinLogging()
	router := gin.New()

	router.LoadHTMLGlob("template/**")
	links = []Link{
		{"title1", "https://a.com"},
		{"title2", "https://b.com"},
	}
	router.GET("/home", Index("list.html"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/home", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "title1/https://a.com/title2/https://b.com/", w.Body.String())
	links = []Link{}
}

func TestNew(t *testing.T) {
	stopGinLogging()
	router := gin.New()

	router.LoadHTMLGlob("template/**")
	router.GET("/urls/new", New("dummy.html"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/urls/new", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())
}

func TestCreate(t *testing.T) {
	stopGinLogging()
	router := gin.New()

	router.LoadHTMLGlob("template/**")
	router.POST("/urls", Create())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/urls", strings.NewReader(`title=title1&url=https%3A%2F%2Fa.com`))
	req.Header = map[string][]string{
		"Content-Type":   {"application/x-www-form-urlencoded"},
		"Content-Length": {"36"},
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, links[0].Title, "title1")
	assert.Equal(t, links[0].Url, "https://a.com")

	assert.Equal(t, 302, w.Code)
}
