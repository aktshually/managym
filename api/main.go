package main

import (
	"os"
	"siege-api/database"
	"siege-api/router"
	"siege-api/utils"

	"github.com/gofiber/fiber/v2"
)

// This API rewrite features user authentication and application management for
// the Siege hosting service.
func main() {
	app := fiber.New()
	utils.LoadEnv()
	database.Connect()

	app.Mount("/", router.MountRoutes())

	app.Listen(os.Getenv("PORT"))
}
