package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		loginId := s.Get("loginId")
		if loginId == nil {
			c.Set("isLogin", false)
		} else {
			_, ok := loginId.(string)
			if ok {
				c.Set("isLogin", true)
			} else {
				c.Set("isLogin", false)
			}
		}
		c.Set("loginId", loginId)
	}
}
