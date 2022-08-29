package session

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store redis.Store
var hasStore = false

func Middleware() gin.HandlerFunc {
	if !hasStore {
		store = getStore()
		hasStore = true
	}
	return sessions.Sessions("url-plumber", store)
}

func getStore() redis.Store {
	addr := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")
	key := os.Getenv("STORE_KEY")
	fmt.Println("addres", addr, password)
	s, err := redis.NewStore(10, "tcp", addr, password, []byte(key))
	if err != nil {
		log.Println(err)
	}
	return s
}
