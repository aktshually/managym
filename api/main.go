package main

import (
	"managym-api/app/users"
	"managym-api/database"
	"managym-api/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

// This API rewrite features user authentication and application management for
// the Managym app.
func main() {
	app := fiber.New()
	utils.LoadEnv()
	database.Connect()

	users.UsersResources{}.MountRoutesInto(app)

	app.Listen(":" + os.Getenv("PORT"))
}
