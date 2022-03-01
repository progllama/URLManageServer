package main

import (
	"url_manager/configs"
	"url_manager/db"
	"url_manager/server"
)

func main() {
	configs.LoadConfig()

	dbms := configs.DBConfig["dbms"]
	dsn := db.BuildDNS(configs.DBConfig)
	db.Open(dbms, dsn)
	db.Migrate()
	defer db.Close()

	server.Open(configs.ServerConfig["port"])
}
