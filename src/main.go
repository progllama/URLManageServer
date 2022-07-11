package main

import (
	"fmt"
	"log"
	"os"
	"url_manager/database"
	"url_manager/server"

	"github.com/joho/godotenv"
)

func main() {
	setLogLevel()
	loadEnvironment()
	database.Open()
	server.Open(":8000")
}

func loadEnvironment() {
	err := godotenv.Load()
	if os.Getenv("MODE") == "dev" {
		vars := []string{
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DBNAME"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("SESSION_NAME"),
			os.Getenv("REDIS_SIZE"),
			os.Getenv("REDIS_NETWORK"),
			os.Getenv("REDIS_ADDRESS"),
			os.Getenv("PASSWORD"),
			os.Getenv("REDIS_SECRET"),
			os.Getenv("REDIRECT_URL"),
			os.Getenv("CRED_FILE_PATH"),
			os.Getenv("GOOGLE_SECRET"),
			os.Getenv("FAVICON_PATH"),
		}
		for v := range vars {
			fmt.Println(vars[v])
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setLogLevel() {
	log.SetFlags(log.Llongfile)
}
