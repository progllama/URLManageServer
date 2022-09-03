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

func TestAccount(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		account := model.Account{
			OpenId: "3uoafdfakb",
		}
		validator := validator.New()
		err := validator.Struct(account)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("bind valid format json", func(t *testing.T) {
		var account model.Account
		jsonData := `{ "openId" : "abcdefg12345"}`
		err := json.Unmarshal([]byte(jsonData), &account)
		if err != nil {
			t.Error(err)
		}
		if account.OpenId != "abcdefg12345" {
			t.Error("fault binding")
		}
	})

	t.Run("bind invalid format json", func(t *testing.T) {
		var account model.Account
		jsonData := `{ "penId" : "abcdefg12345"}`
		err := json.Unmarshal([]byte(jsonData), &account)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("blank openId should be error", func(t *testing.T) {
		account := model.Account{}
		validator := validator.New()
		err := validator.Struct(account)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("openId length 256 should be no error(validator)", func(t *testing.T) {
		testId := make([]rune, 256)
		for i := 0; i < 256; i++ {
			testId[i] = 'a'
		}
		account := model.Account{
			OpenId: string(testId),
		}
		validator := validator.New()
		err := validator.Struct(account)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("openId longer than 256 should be error(validator)", func(t *testing.T) {
		testId := make([]rune, 257)
		for i := 0; i < 257; i++ {
			testId[i] = 'a'
		}
		account := model.Account{
			OpenId: string(testId),
		}
		validator := validator.New()
		err := validator.Struct(account)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("duplicate openid should be error", func(t *testing.T) {
		account := model.Account{}
		validator := validator.New()
		err := validator.Struct(account)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("openId length 256 should be no error(gorm)", func(t *testing.T) {
		testId := make([]rune, 256)
		for i := 0; i < 256; i++ {
			testId[i] = 'a'
		}
		account := model.Account{OpenId: string(testId)}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.Account{})
		err := db.Create(&account).Error
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("openId length 257 should be error(gorm)", func(t *testing.T) {
		testId := make([]rune, 257)
		for i := 0; i < 257; i++ {
			testId[i] = 'a'
		}
		account := model.Account{OpenId: string(testId)}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.Account{})
		db.Create(&account)
		err := db.Create(&account).Error
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("blank openId should be error", func(t *testing.T) {
		account := model.Account{}
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		con, _ := db.DB()
		defer con.Close()
		db.AutoMigrate(model.Account{})
		err := db.Create(&account).Error
		if err == nil {
			t.Error(err)
		}
	})
}
