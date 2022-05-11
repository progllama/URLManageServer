package db

import (
	"url_manager/app/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func OpenSqlite() {
	db, err = gorm.Open(sqlite.Open("test.db"))
}

func Open(database string, dsn string) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Close() {
	connection, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = connection.Close()
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Url{})
}
