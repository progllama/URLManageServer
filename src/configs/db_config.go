package configs

import (
	"os"
)

var DBConfig = make(map[string]string)

func loadDBConfig() {
	DBConfig["dbms"] = os.Getenv("DBMS")
	DBConfig["host"] = os.Getenv("HOST")
	DBConfig["port"] = os.Getenv("PORT")
	DBConfig["user"] = os.Getenv("USER")
	DBConfig["dbname"] = os.Getenv("DBNAME")
	DBConfig["password"] = os.Getenv("PASSWORD")
}
