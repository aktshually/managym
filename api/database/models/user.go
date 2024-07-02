// Package `models` provides the base models for the application
package models

// This struct will be used to generate users for the database
type User struct {
	Id        string `gorm:"primaryKey"`
	Email     string `gorm:"<-;size:120;index"`
	Password  string `gorm:"<-;size:120"`
	FirstName string `gorm:"size:120"`
	LastName  string `gorm:"size:120"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	IsActive  bool   `gorm:"default:true"`
}
