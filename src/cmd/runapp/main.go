package main

import (
	"log"
	"url_manager/configs"
	"url_manager/db"
	"url_manager/server"
)

func main() {
	log.SetFlags(log.Llongfile)

	configs.LoadConfig()

	// dbms := configs.DBConfig["dbms"]
	// dsn := db.BuildDNS(configs.DBConfig)
	// db.Open(dbms, dsn)
	db.OpenSqlite()
	db.Migrate()
	defer db.Close()

	server.Open(configs.ServerConfig["port"])
}
