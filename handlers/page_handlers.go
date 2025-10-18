package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// PageHandler handles page rendering
type PageHandler struct{}

// NewPageHandler creates a new PageHandler
func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

// Index renders the home page
func (h *PageHandler) Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":   "Welcome to Fiber",
		"Message": "This is a Go Fiber API with HTML templates!",
	})
}

// About renders the about page
func (h *PageHandler) About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title":   "About",
		"Content": "This is a demonstration of Fiber with HTML templating and HTMX.",
	})
}
