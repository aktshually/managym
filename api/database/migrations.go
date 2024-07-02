package database

import (
	"siege-api/database/models"

	"gorm.io/gorm"
)

// Generates migrations for all the models listed in `database/models`
func GenerateMigrations(tx *gorm.DB) {
	tx.AutoMigrate(&models.User{})
}
