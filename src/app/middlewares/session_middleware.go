package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("UserId")

		if uid == nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
		} else {
			c.Set("UserId", uid)
			c.Next()
		}
	}
}
