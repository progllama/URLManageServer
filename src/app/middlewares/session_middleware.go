package middlewares

import (
	"log"
	"net/http"
	"url_manager/app/session"

	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := session.NewRedisSession(c)

		if session.HasUserId() {
			c.Next()
			return
		}

		log.Println(" please login.")
		c.Redirect(http.StatusFound, "/users")
		c.Abort()
	}
}
