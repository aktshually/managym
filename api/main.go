package main

import (
	"managym/app/users"
	"managym/database"
	"managym/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// This API rewrite features user authentication and application management for
// the Managym app.
func main() {
	app := fiber.New()
	utils.LoadEnv()
	database.Connect()

	app.Use(cors.New())

	users.UsersResources{}.MountRoutesInto(app)

	app.Listen(":" + os.Getenv("PORT"))
}
