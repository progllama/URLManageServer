package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")
		if uid == nil {
			c.Set("loggsin", false)
			c.Redirect(302, "/about")
			c.Abort()
		} else {
			c.Set("logsin", true)
			c.Next()
		}
	}
}
