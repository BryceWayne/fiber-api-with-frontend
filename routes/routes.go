package routes

import (
	"github.com/BryceWayne/fiber-api-with-frontend/handlers"
	"github.com/gofiber/fiber/v2"
)

// Setup configures all application routes
func Setup(app *fiber.App) {
	// Initialize handlers
	pageHandler := handlers.NewPageHandler()
	apiHandler := handlers.NewAPIHandler()

	// Serve static files
	app.Static("/static", "./static")

	// Page routes
	app.Get("/", pageHandler.Index)
	app.Get("/about", pageHandler.About)

	// API routes
	api := app.Group("/api")
	api.Get("/hello", apiHandler.Hello)
	api.Get("/hello-html", apiHandler.HelloHTML)
}
