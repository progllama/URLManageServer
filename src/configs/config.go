package configs

import (
	"github.com/joho/godotenv"
)

func LoadConfig() {
	loadDotEnv()
	loadDBConfig()
	loadRedisConfig()
	loadServerConfig()
}

func loadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
