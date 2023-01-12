package router

import (
	"github.com/bmdavis419/the-better-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// setup health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// setup the todos group
	todos := app.Group("/todos")
	todos.Get("/", handlers.HandleAllTodos)
	todos.Post("/", handlers.HandleCreateTodo)
	todos.Put("/:id", handlers.HandleUpdateTodo)
	todos.Get("/:id", handlers.HandleGetOneTodo)
	todos.Delete("/:id", handlers.HandleDeleteTodo)
}
