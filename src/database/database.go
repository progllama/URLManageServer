package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	"url_manager/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := getDSN()
	config := getConfig()
	var err error
	db, err = openDB(dsn, config)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Link{},
	)
}

func Disconnect() {
	if con, err := db.DB(); err != nil {
		con.Close()
	}
}

func GetDB() *gorm.DB {
	return db
}

func getDSN() string {
	keys := dsnKeys()
	dsn := make([]string, len(keys))
	for i, key := range keys {
		value := os.Getenv(strings.ToUpper(key))
		dsn[i] = fmt.Sprintf("%s=%s ", key, value)
	}
	return strings.Join(dsn, " ")
}

func getConfig() *gorm.Config {
	return &gorm.Config{}
}

func openDB(dsn string, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), config)
}

func dsnKeys() []string {
	return []string{
		"host",
		"user",
		"password",
		"dbname",
		"port",
		"sslmode",
		"TimeZone",
	}
}
