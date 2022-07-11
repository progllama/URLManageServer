package session

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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
	store := memstore.NewStore([]byte(secret))
	return store
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
