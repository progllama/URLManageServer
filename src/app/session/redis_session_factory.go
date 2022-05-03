package session

import "github.com/gin-gonic/gin"

type RedisSessionFactory struct {
}

func (factory *RedisSessionFactory) Create(c *gin.Context) Session {
	return NewRedisSession(c)
}
