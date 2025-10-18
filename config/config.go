package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// DatabaseType represents the type of database to use
type DatabaseType string

const (
	DatabaseMemory   DatabaseType = "memory"
	DatabaseFirebase DatabaseType = "firebase"
)

// Config holds application configuration
type Config struct {
	Port             string
	ViewsPath        string
	StaticPath       string
	TemplateExt      string
	DatabaseType     DatabaseType
	FirebaseProject  string
	FirebaseCollection string
}

// NewConfig creates a new configuration with default values
// Database configuration can be overridden via environment variables:
// - DB_TYPE: "memory" or "firebase" (default: "memory")
// - FIREBASE_PROJECT_ID: Google Cloud project ID
// - FIREBASE_COLLECTION: Firestore collection name (default: "books")
func NewConfig() *Config {
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "memory"
	}

	firebaseProject := os.Getenv("FIREBASE_PROJECT_ID")
	firebaseCollection := os.Getenv("FIREBASE_COLLECTION")
	if firebaseCollection == "" {
		firebaseCollection = "books"
	}

	return &Config{
		Port:               ":3000",
		ViewsPath:          "./views",
		StaticPath:         "./static",
		TemplateExt:        ".html",
		DatabaseType:       DatabaseType(dbType),
		FirebaseProject:    firebaseProject,
		FirebaseCollection: firebaseCollection,
	}
}

// SetupTemplateEngine initializes and returns the HTML template engine
func (c *Config) SetupTemplateEngine() *html.Engine {
	return html.New(c.ViewsPath, c.TemplateExt)
}

// SetupApp creates and configures a new Fiber app
func (c *Config) SetupApp(engine *html.Engine) *fiber.App {
	return fiber.New(fiber.Config{
		Views: engine,
	})
}
