package db

import (
	"fmt"
	"url_manager/models"

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

func BuildDNS(params map[string]string) string {
	isValid := validate(params)
	if !isValid {
		panic("DNS parameter is not valid.")
	} else {
		return build(params)
	}
}

func validate(params map[string]string) bool {
	return true
}

func build(params map[string]string) string {
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		params["host"],
		params["port"],
		params["user"],
		params["dbname"],
		params["password"],
	)
	fmt.Println(dns)
	return dns
}
