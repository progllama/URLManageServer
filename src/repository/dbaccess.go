package repository

import (
	"url_manager/database"

	"gorm.io/gorm"
)

func getDB() *gorm.DB {
	return database.GetDB()
}
