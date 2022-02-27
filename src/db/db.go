package db

import (
	"url_manager/utils"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func Open(database string, dsn string) {
	db, err = gorm.Open(database, dsn)
	utils.PanicIfError(err)
}

func Close() {
	err = db.Close()
	utils.PanicIfError(err)
}
