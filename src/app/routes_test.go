package app

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoutes(t *testing.T) {
	stopGinLogging()
	Routes(gin.Default())
}
