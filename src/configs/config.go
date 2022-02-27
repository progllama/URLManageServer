package configs

import (
	"url_manager/utils"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	loadDotEnv()
	loadDBConfig()
	loadServerConfig()
}

func loadDotEnv() {
	err := godotenv.Load(".env")
	utils.PanicIfError(err)
}
