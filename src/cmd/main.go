package main

import (
	"url_manager/database"
	"url_manager/dotenv"
	"url_manager/server"
)

func main() {
	dotenv.Load()

	database.Connect()
	defer database.Disconnect()

	server.Run()
}
