package initializers

import "github.com/nidhey27/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
