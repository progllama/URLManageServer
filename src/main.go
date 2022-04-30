package main

import (
	"log"
	"time"
	"url_manager/configs"
	"url_manager/db"
	"url_manager/server"
)

func main() {
	time.Sleep(10 * time.Second)

	log.SetFlags(log.Llongfile)

	configs.LoadConfig()

	dbms := configs.DBConfig["dbms"]
	dsn := db.BuildDNS(configs.DBConfig)
	db.Open(dbms, dsn)
	db.Migrate()
	defer db.Close()

	server.Open(configs.ServerConfig["port"])
}
