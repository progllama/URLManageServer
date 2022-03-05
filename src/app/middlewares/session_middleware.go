package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")
		if uid == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "please login."})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
