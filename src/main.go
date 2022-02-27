package main

import (
	"url_manager/configs"
	"url_manager/db"
	"url_manager/server"
)

func main() {
	configs.LoadConfig()

	dbsm := configs.Config["dbsm"]
	dsn := db.BuildDNS(configs.Config)
	db.Open(dbsm, dsn)
	defer db.Close()

	server.Open(configs.ServerConfig["port"])
}
