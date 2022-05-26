package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// apiのテストならレスポンスも確かめたほうがいいかも?
func TestNewUserSuccess(t *testing.T) {
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
	// ctrl := New
	// r.GET(url, ctrl.NewSession)

	// call
	r.ServeHTTP(w, req)

	// check status code.
	assert.Equal(t, 200, w.Code)
}
