// Package `utils` implements utilities that will be useful through the development stage
package utils

import "github.com/joho/godotenv"


// Loads environment variables from the ".env" file.
// As the application depends on the information provided by the `godotenv.Load` method,
// if the file can not be loaded, it will panic and the server will be shut down.
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
