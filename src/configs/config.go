package configs

import (
	"url_manager/utils"

	"github.com/joho/godotenv"
)

var Config map[string]string = make(map[string]string)

func LoadConfig() {
	loadDotEnv()
	loadDBConfig()
	loadServerConfig()
}

func loadDotEnv() {
	err := godotenv.Load()
	utils.PanicIfError(err)
}
