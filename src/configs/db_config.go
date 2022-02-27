package configs

import (
	"os"
)

var DBConfig = make(map[string]string)

func loadDBConfig() {
	DBConfig["dbsm"] = os.Getenv("DBSM")
	DBConfig["host"] = os.Getenv("HOST")
	DBConfig["port"] = os.Getenv("PORT")
	DBConfig["user"] = os.Getenv("USER")
	DBConfig["dbname"] = os.Getenv("DBNAME")
	DBConfig["password"] = os.Getenv("PASS")
}
