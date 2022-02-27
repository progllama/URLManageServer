package main

import (
	"fmt"
	"url_manager/configs"
	"url_manager/db"
	"url_manager/server"
)

func main() {
	configs.LoadConfig()

	dbsm := configs.DBConfig["dbms"]
	dsn := db.BuildDNS(configs.DBConfig)
	fmt.Println(dsn)
	db.Open(dbsm, dsn)
	defer db.Close()

	server.Open(configs.ServerConfig["port"])
}
