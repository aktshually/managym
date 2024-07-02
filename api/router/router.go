// Package `router` provides features that will help to deal with routing.
package router

import "github.com/gofiber/fiber/v2"

// Joins all the routes into a new router, that will be mounted into the main app during
// server initialization.
func MountRoutes() *fiber.App {
	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hey")
	})

	return router
}
