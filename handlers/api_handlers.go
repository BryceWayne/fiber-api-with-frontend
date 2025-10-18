package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// APIHandler handles API endpoints
type APIHandler struct{}

// NewAPIHandler creates a new APIHandler
func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

// Hello returns a hello message
func (h *APIHandler) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello from Fiber API!",
	})
}

// HelloHTML returns an HTML fragment for HTMX
func (h *APIHandler) HelloHTML(c *fiber.Ctx) error {
	return c.SendString(`<p class="api-result">API Response: Hello from Fiber API!</p>`)
}
