package routes

import (
	"context"
	"fmt"
	"log"

	"github.com/BryceWayne/fiber-api-with-frontend/config"
	_ "github.com/BryceWayne/fiber-api-with-frontend/docs" // Import generated docs
	"github.com/BryceWayne/fiber-api-with-frontend/handlers"
	"github.com/BryceWayne/fiber-api-with-frontend/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// initRepository initializes the repository based on configuration
func initRepository(cfg *config.Config) (repository.BookRepository, error) {
	switch cfg.DatabaseType {
	case config.DatabaseFirebase:
		log.Printf("Initializing Firebase repository with project: %s, collection: %s", 
			cfg.FirebaseProject, cfg.FirebaseCollection)
		return repository.NewFirebaseRepository(
			context.Background(),
			cfg.FirebaseProject,
			cfg.FirebaseCollection,
		)
	case config.DatabaseMemory:
		fallthrough
	default:
		log.Println("Initializing in-memory repository")
		return repository.NewMemoryRepository(), nil
	}
}

// Setup configures all application routes
func Setup(app *fiber.App, cfg *config.Config) error {
	// Initialize repository based on configuration
	repo, err := initRepository(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize repository: %w", err)
	}

	// Initialize handlers
	pageHandler := handlers.NewPageHandler()
	apiHandler := handlers.NewAPIHandler(repo)

	// Serve static files
	app.Static("/static", "./static")

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Page routes
	app.Get("/", pageHandler.Index)
	app.Get("/about", pageHandler.About)

	// API routes
	api := app.Group("/api")
	api.Get("/hello", apiHandler.Hello)
	api.Get("/hello-html", apiHandler.HelloHTML)

	// Book CRUD routes
	api.Post("/books", apiHandler.CreateBook)
	api.Get("/books", apiHandler.GetBooks)
	api.Get("/books/:id", apiHandler.GetBook)
	api.Put("/books/:id", apiHandler.UpdateBook)
	api.Delete("/books/:id", apiHandler.DeleteBook)
	
	return nil
}
