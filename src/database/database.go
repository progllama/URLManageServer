package database

import (
	"fmt"
	"os"
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

func Database() *gorm.DB {
	return db
}

func Open(database string, dsn string) {
	if os.Getenv("MODE") == "dev" {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
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
	db.AutoMigrate(&models.Link{})
	db.AutoMigrate(&models.LinkList{})
	db.AutoMigrate(&models.LinkListRelation{})
}

func BuildDNS(host, port, user, dbname, password string) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)
}
