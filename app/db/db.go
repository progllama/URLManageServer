package db

import (
	"fmt"
	"os"
	"url_manager/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	err := godotenv.Load(".env") // envファイルのパスを渡す。何も渡さないと、どうディレクトリにある、.envファイルを探す
	if err != nil {
		panic("Error loading .env file")
	}

	dbms := os.Getenv("DBMS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASS")

	db, err = gorm.Open(dbms, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password))
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&model.User{})
}
