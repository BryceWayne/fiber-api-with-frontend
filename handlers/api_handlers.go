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

// HTML response template for HTMX
const helloHTMLTemplate = `<p class="api-result">API Response: Hello from Fiber API!</p>`

// HelloHTML returns an HTML fragment for HTMX
func (h *APIHandler) HelloHTML(c *fiber.Ctx) error {
	return c.SendString(helloHTMLTemplate)
}
