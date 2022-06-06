package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const CSRF_SECRET = "CSRF_SECRET"

// 事前にSessionが登録されている必要あり。
func RegisterCsrfMiddleware(router *gin.Engine) {
	secret := os.Getenv(CSRF_SECRET)
	router.Use(csrf.Middleware(csrf.Options{
		Secret: secret,
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.Abort()
		},
	}))
}
