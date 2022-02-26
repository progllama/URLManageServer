package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func Open() {
	config := loadConfig()
	dsnParams = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.host,
		config.port,
		config.user,
		config.dbname,
		password
	)

	db, err = gorm.Open(
		config.Database,
		dsnParams)
	if err != nil {
		panic(err)
	}
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func getDSNParametors() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.host,
		config.port,
		config.user,
		config.dbname,
		config.password
	)
}
