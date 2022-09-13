package main

import (
	"log"
	"os"
	"url_manager/auth"
	"url_manager/dotenv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dotenv.Load()
	log.Println(os.Getenv("GOOGLE_CLIENT_ID"))
	auth.GoogleConf.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	auth.GoogleConf.ClientSecret = os.Getenv("GOOGLE_SECRET")

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	auth.DB = db
	db.AutoMigrate(&auth.ConfidentialAccount{})

	engine := gin.Default()

	store := memstore.NewStore([]byte("secret"))
	router := engine.Use(sessions.Sessions("test", store))

	router.GET("/", func(ctx *gin.Context) {
		state := auth.RandomState()
		session := sessions.Default(ctx)
		session.Set("state", state)
		session.Save()
		v, _ := auth.CreateCodeVerifier()
		codeChallenge := v.CodeChallengeS256()
		session.Set("verifier", v.String())
		session.Save()
		ctx.Writer.Write([]byte(`
	<html>
		<head>
			<title>Golang Google</title>
		</head>
		<body>
			<a href='` + auth.GoogleLoginURL(state, codeChallenge) + `'>
				<button>Login with Facebook!</button>
			</a>
		</body>
	</html>`))
	})
	router.GET("/login", auth.GoogleOAuth2RedirectHandler)

	engine.Run(":8080")
}
