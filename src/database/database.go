package database

import (
	"fmt"
	"log"
	"url_manager/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	HOST     string
	USER     string
	PASSWORD string
	DBNAME   string
	PORT     string
	SSLMODE  string
	TIMEZONE string
)

func Connect() {
	dsn := DSN()
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(
		&model.User{},
		&model.Link{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect() {
	if con, err := DB.DB(); err != nil {
		con.Close()
	}
}

func DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		HOST, USER, PASSWORD, DBNAME, PORT, SSLMODE, TIMEZONE,
	)
}

func openDB(dsn string, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
