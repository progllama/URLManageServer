package db

import "os"

type Config struct {
	DB       string
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func loadConfig() Config {
	loadEnv()
	return Config{
		os.Getenv("DB"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
	}
}
