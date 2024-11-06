package database

import (
	"fmt"

	"github.com/anishNagula/pesuio-final-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize the DB connection (run this at the start of your app)
func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("database connection failed: %v", err)
	}

	// Migrate the schema for models.User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to migrate schema: %v", err)
	}

	return nil
}

// Create a new user
func CreateUser(username, password string) error {
	user := models.User{Username: username, Password: password}

	var existingUser models.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return fmt.Errorf("user with username %s already exists", username)
	}

	result := db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %v", result.Error)
	}

	return nil
}

func CheckPassword(username, password string) (bool, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return false, fmt.Errorf("user not found: %v", err)
	}

	if user.Password == password {
		return true, nil
	}

	return false, nil
}
