package session

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var (
	sessionName string
	size        int
	network     string
	address     string
	password    string
	secret      string
	err         error
)

func NewSessionHandler() gin.HandlerFunc {
	loadConfig()

	store := makeStore()
	return sessions.Sessions(sessionName, store)
}

func makeStore() sessions.Store {
	if os.Getenv("MODE") == "dev" {
		store := cookie.NewStore([]byte("on develop"))
		if err != nil {
			log.Fatal(err)

		}
		return store
	} else {
		store, err := redis.NewStore(size, network, address, password, []byte(secret))
		if err != nil {
			log.Fatal(err)
		}
		return store
	}
}

func loadConfig() {
	sessionName = os.Getenv("SESSION_NAME")

	size, err = strconv.Atoi(os.Getenv("REDIS_SIZE"))
	if err != nil {
		log.Fatal(err)
	}

	network = os.Getenv("REDIS_NETWORK")
	address = os.Getenv("REDIS_ADDRESS")
	password = os.Getenv("PASSWORD")
	secret = os.Getenv("REDIS_SECRET")
}
