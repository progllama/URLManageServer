package main

import (
	"log"
	"os"
	"url_manager/auth"
	controllers "url_manager/controller"
	"url_manager/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	auth.GoogleConf.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	auth.GoogleConf.ClientSecret = os.Getenv("GOOGLE_SECRET")
	auth.GoogleConf.RedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")
	auth.FacebookConf.ClientID = os.Getenv("FACEBOOK_CLIENT_ID")
	auth.FacebookConf.ClientSecret = os.Getenv("FACEBOOK_SECRET")
	auth.FacebookConf.RedirectURL = os.Getenv("FACEBOOK_REDIRECT_URL")
	auth.TwitterConf.ClientID = os.Getenv("TWITTER_CLIENT_ID")
	auth.TwitterConf.ClientSecret = os.Getenv("TWITTER_SECRET")
	auth.TwitterConf.RedirectURL = os.Getenv("TWITTER_REDIRECT_URL")

	database.HOST = os.Getenv("HOST")
	database.USER = os.Getenv("USER")
	database.PASSWORD = os.Getenv("PASSWORD")
	database.DBNAME = os.Getenv("DBNAME")
	database.PORT = os.Getenv("PORT")
	database.SSLMODE = os.Getenv("SSLMODE")
	database.TIMEZONE = os.Getenv("TIMEZONE")

	addr := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")
	key := os.Getenv("STORE_KEY")

	database.Connect()
	defer database.Disconnect()

	e := gin.Default()
	e.LoadHTMLGlob("public/index.html")
	e.Static("/static", "./public/static")
	s, err := redis.NewStore(10, "tcp", addr, password, []byte(key))
	if err != nil {
		log.Fatal(err)
	}

	router := e.Use(sessions.Sessions("cache", s))
	router.GET("/", controllers.Entry)
	router.GET("/login/google", auth.GoogleAuth)
	router.GET("/login/url/google", auth.GetGoogleLoginURL)
	router.GET("/login/facebook", auth.FacebookAuth)
	router.GET("/login/url/facebook", auth.GetFacebookLoginURL)
	router.GET("/login/twitter", auth.TwitterAuth)
	router.GET("/login/url/twitter", auth.GetTwitterLoginURL)

	router.GET("/users/me", controllers.GetMe)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
