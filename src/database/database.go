package database

import (
	"fmt"
	"log"
	"url_manager/models"

	_ "github.com/lib/pq"
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

func Open() {
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Migrate()
}

func Close() {
	connection, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Close()
	if err != nil {
		log.Fatal(err)
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
