package db

import (
	"url_manager/app/controllers"
)

func Migrate() {
	db.AutoMigrate(&model.User{})
}
