package session

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func NewSessionHandler() gin.HandlerFunc {
	// get configs.
	sessionName := os.Getenv("SESSION_NAME")
	size := os.Getenv("REDIS_SIZE")
	network := os.Getenv("REDIS_NETWORK")
	address := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("PASSWORD")
	secret := os.Getenv("REDIS_SECRET")

	// type conversion.
	intSize, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal(err)
	}

	// make store.
	store, err := redis.NewStore(intSize, network, address, password, []byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return sessions.Sessions(sessionName, store)
}
