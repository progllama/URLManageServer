package model_test

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// please do yourself that migrate and close process
func createTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return db, err
}

// test utility.
// apply validator and return it's result
func valid(s interface{}) (err error) {
	validator := validator.New()
	err = validator.Struct(s)
	return
}

// test utility.
// return long(256) name
func getLongName() string {
	return "alicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicealicea"
}
