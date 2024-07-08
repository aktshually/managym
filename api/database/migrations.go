package database

import (
	"managym-api/database/models"

	"gorm.io/gorm"
)

// Generates migrations for all the models listed in `database/models`. It will
// panic in case the migrations can not be generated, because they are essential
// for the application to work.
func GenerateMigrations(tx *gorm.DB) {
	err := tx.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
