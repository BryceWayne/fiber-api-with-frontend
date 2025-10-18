package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber instance with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	app.Static("/static", "./static")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":   "Welcome to Fiber",
			"Message": "This is a Go Fiber API with HTML templates!",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		// Render about template
		return c.Render("about", fiber.Map{
			"Title":   "About",
			"Content": "This is a demonstration of Fiber with HTML templating.",
		})
	})

	// API endpoint
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from Fiber API!",
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
