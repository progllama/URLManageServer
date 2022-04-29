package configs

import (
	"log"
	"os"
)

var DBConfig = make(map[string]string)

func loadDBConfig() {
	log.Println(os.Getenv("DBMS"))
	DBConfig["dbms"] = os.Getenv("DBMS")
	DBConfig["host"] = os.Getenv("HOST")
	DBConfig["port"] = os.Getenv("PORT")
	DBConfig["user"] = os.Getenv("USER")
	DBConfig["dbname"] = os.Getenv("DBNAME")
	DBConfig["password"] = os.Getenv("PASSWORD")
}
