package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Config holds application configuration
type Config struct {
	Port        string
	ViewsPath   string
	StaticPath  string
	TemplateExt string
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	return &Config{
		Port:        ":3000",
		ViewsPath:   "./views",
		StaticPath:  "./static",
		TemplateExt: ".html",
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
