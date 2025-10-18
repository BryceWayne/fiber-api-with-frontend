package main

import (
	"log"

	"github.com/BryceWayne/fiber-api-with-frontend/config"
	"github.com/BryceWayne/fiber-api-with-frontend/routes"
)

// @title Fiber API with Frontend
// @version 1.0
// @description A modular web application built with Go Fiber framework featuring RESTful API endpoints, HTML templating, and HTMX integration.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/BryceWayne/fiber-api-with-frontend
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @schemes http

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Setup template engine
	engine := cfg.SetupTemplateEngine()

	// Create Fiber app
	app := cfg.SetupApp(engine)

	// Setup routes
	routes.Setup(app, cfg)

	// Start server
	log.Fatal(app.Listen(cfg.Port))
}
