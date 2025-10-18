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
		"Title":   "Lorem Ipsum Dolor Sit Amet",
		"Message": "Consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua!",
	})
}

// About renders the about page
func (h *PageHandler) About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title":   "Ut Enim Ad Minim",
		"Content": "Veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
	})
}
