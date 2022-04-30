package db

import (
	"time"
	"url_manager/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func Open(database string, dsn string) {
	time.Sleep(10 * time.Second)
	db, err = gorm.Open(database, dsn)
	if err != nil {
		panic(err)
	}
}

func Close() {
	err = db.Close()
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Url{})
}
