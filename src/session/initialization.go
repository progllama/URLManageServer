package session

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func InitializeSession(router *gin.Engine) {
	sessionType := os.Getenv(SESSION_TYPE)
	secret := os.Getenv(SESSION_SECRET)

	var store sessions.Store = memstore.NewStore([]byte(secret))
	switch sessionType {
	case COOKIE:
		store = cookie.NewStore([]byte(secret))
	case REDIS:
	case MEM_CACHED:
	case MONGO_DB:
	case GORM:
	case MEMORY:
		store = memstore.NewStore([]byte(secret))
	case POSTGRESQL:
	default:
		log.Println("used invalid session type.")
	}
	router.Use(sessions.Sessions("mysession", store))
}
