package main

import (
	"log"

	"github.com/BryceWayne/fiber-api-with-frontend/config"
	"github.com/BryceWayne/fiber-api-with-frontend/routes"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Setup template engine
	engine := cfg.SetupTemplateEngine()

	// Create Fiber app
	app := cfg.SetupApp(engine)

	// Setup routes
	routes.Setup(app)

	// Start server
	log.Fatal(app.Listen(cfg.Port))
}
