package session

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store = getStore()

func Middleware() gin.HandlerFunc {
	return sessions.Sessions("url-plumber", store)
}

func getStore() redis.Store {
	addr := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")
	key := os.Getenv("STORE_KEY")
	s, err := redis.NewStore(10, "tcp", addr, password, []byte(key))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
