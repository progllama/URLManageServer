package app

import (
	"io"

	"github.com/gin-gonic/gin"
)

func stopGinLogging() {
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
}
