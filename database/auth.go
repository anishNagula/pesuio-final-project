package database

import (
	"fmt"

	"github.com/anishNagula/pesuio-final-project/models"
)

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
