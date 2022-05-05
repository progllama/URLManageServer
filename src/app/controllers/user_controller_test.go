package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// apiのテストならレスポンスも確かめたほうがいいかも?
func TestNewUser(t *testing.T) {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}

	ctx.Request = req
	asserts := assert.New(t)

	ctrl := NewUserController(nil, nil)
	ctrl.New(ctx)

	asserts.Equal(t, 200, rec.Code)
}
