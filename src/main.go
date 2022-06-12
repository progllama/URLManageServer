package main

import (
	"log"
	"url_manager/server"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Llongfile)
	loadEnv()
	server.Open(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
