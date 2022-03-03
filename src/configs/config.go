package configs

import (
	"github.com/joho/godotenv"
)

func LoadConfig() {
	loadDotEnv()
	loadDBConfig()
	loadServerConfig()
}

func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}
