// Package `database` will implement database related methods
package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Per the Gorm documentation, a new connection pool is created for each `gorm.DB` object,
// so it is recommended to call this method only once and reuse the returned `gorm.DB` object.
var DatabaseManager *gorm.DB

// Connects to a PostgreSQL database and checks whether the connection ran successfully or not.
// As the database connection is absolutely necessary so that the application is fully functional,
// the server will be shut down if the connection is not successfully made.
func Connect() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_HOST"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	GenerateMigrations(db)

	DatabaseManager = db
}

func MockConnection() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	GenerateMigrations(db)

	DatabaseManager = db
}
