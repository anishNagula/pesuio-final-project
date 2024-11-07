package database

import (
	"fmt"

	"github.com/anishNagula/pesuio-final-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("leet.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Question{}, &models.TestCase{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
