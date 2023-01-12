package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// setup health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
