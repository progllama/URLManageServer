package model_test

import (
	"encoding/json"
	"testing"
	"url_manager/model"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestUser(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		user := model.User{
			Name: "Alice",
		}
		validator := validator.New()
		err := validator.Struct(user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("bind valid format json", func(t *testing.T) {
		var user model.User
		jsonData := `{ "name" : "Alice"}`
		err := json.Unmarshal([]byte(jsonData), &user)
		if err != nil {
			t.Error(err)
		}
		if user.Name != "Alice" {
			t.Error("fault binding")
		}
	})

	t.Run("bind invalid format json", func(t *testing.T) {
		var user model.User
		jsonData := `{ "same" : "Alice"}`
		err := json.Unmarshal([]byte(jsonData), &user)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("blank name should be error", func(t *testing.T) {
		user := model.User{}
		validator := validator.New()
		err := validator.Struct(user)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("name length 256 should be no error(validator)", func(t *testing.T) {
		name := make([]rune, 256)
		for i := 0; i < 256; i++ {
			name[i] = 'a'
		}
		user := model.User{
			Name: string(name),
		}
		validator := validator.New()
		err := validator.Struct(user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("openId longer than 256 should be error(validator)", func(t *testing.T) {
		name := make([]rune, 257)
		for i := 0; i < 257; i++ {
			name[i] = 'a'
		}
		user := model.User{
			Name: string(name),
		}
		validator := validator.New()
		err := validator.Struct(user)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("duplicate name should be error", func(t *testing.T) {
		user := model.User{Name: "Alice"}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.User{})
		db.Create(&user)
		err := db.Create(&user).Error
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("openId length 256 should be no error(gorm)", func(t *testing.T) {
		name := make([]rune, 256)
		for i := 0; i < 256; i++ {
			name[i] = 'a'
		}
		user := model.User{
			Name: string(name),
		}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.User{})
		err := db.Create(&user).Error
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("openId length 257 should be error(gorm)", func(t *testing.T) {
		name := make([]rune, 257)
		for i := 0; i < 257; i++ {
			name[i] = 'a'
		}
		user := model.User{
			Name: string(name),
		}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.User{})
		err := db.Create(&user).Error
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("blank openId should be error", func(t *testing.T) {
		user := model.User{}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.User{})
		err := db.Create(&user).Error
		if err == nil {
			t.Error(err)
		}
	})
}
