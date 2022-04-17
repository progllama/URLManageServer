package middlewares

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println()
		session := sessions.Default(c)
		uid := session.Get("login_id")
		if uid == nil {
			c.Set("loggedin", false)
			c.Redirect(302, "/home")
			c.Abort()
		} else {
			c.Set("loggedin", true)
			c.Next()
		}
	}
}
