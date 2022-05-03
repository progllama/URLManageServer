package session

import "github.com/gin-gonic/gin"

type SessionFactory interface {
	Create(c *gin.Context) Session
}
