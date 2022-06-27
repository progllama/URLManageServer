package main

import (
	"log"
	"url_manager/server"

	"github.com/joho/godotenv"
)

func main() {
	setLogLevel()
	loadEnvironment()
	server.Open("PORT")
}

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setLogLevel() {
	log.SetFlags(log.Llongfile)
}
