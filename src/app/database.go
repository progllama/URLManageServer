package app

import (
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (database *Database) Open() error {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	database.db = db
	db.AutoMigrate(&Link{})
	return nil
}

func (database *Database) Close() error {
	db, err := database.db.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	return err
}
