package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
}

func DestroySession(c *gin.Context) {
	c.Redirect(302, "/users")
}

func SessionCheck(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("UserId")

	if uid == nil {
		c.Redirect(302, "/sing_in")
		c.Abort()
	} else {
		c.Set("UserId", uid)
		c.Next()
	}
}
