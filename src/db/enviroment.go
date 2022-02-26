package db

import "github.com/joho/godotenv"

const envFilePath = ".env"

func loadEnv() {
	// Argument is .env file path. if it doesn't exist, use .env in same directory.
	err := godotenv.Load(envFilePath)
	if err != nil {
		panic("Fail to load database configration.")
	}
}
