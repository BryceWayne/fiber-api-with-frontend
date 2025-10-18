package handlers

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Book represents a book in the bookstore
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Year   int    `json:"year"`
}

// APIHandler handles API endpoints
type APIHandler struct {
	books      []Book
	booksMutex sync.RWMutex
	nextID     int
}

// NewAPIHandler creates a new APIHandler
func NewAPIHandler() *APIHandler {
	return &APIHandler{
		books:  []Book{},
		nextID: 1,
	}
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

// CreateBook handles POST /api/books - creates a new book
func (h *APIHandler) CreateBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if book.Title == "" || book.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and Author are required fields",
		})
	}

	h.booksMutex.Lock()
	book.ID = h.nextID
	h.nextID++
	h.books = append(h.books, *book)
	h.booksMutex.Unlock()

	return c.Status(fiber.StatusCreated).JSON(book)
}

// GetBooks handles GET /api/books - returns all books
func (h *APIHandler) GetBooks(c *fiber.Ctx) error {
	h.booksMutex.RLock()
	defer h.booksMutex.RUnlock()

	return c.JSON(h.books)
}

// GetBook handles GET /api/books/:id - returns a specific book
func (h *APIHandler) GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	h.booksMutex.RLock()
	defer h.booksMutex.RUnlock()

	for _, book := range h.books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Book not found",
	})
}

// UpdateBook handles PUT /api/books/:id - updates an existing book
func (h *APIHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	updatedBook := new(Book)
	if err := c.BodyParser(updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	h.booksMutex.Lock()
	defer h.booksMutex.Unlock()

	for i, book := range h.books {
		if book.ID == id {
			// Preserve the ID
			updatedBook.ID = id
			h.books[i] = *updatedBook
			return c.JSON(updatedBook)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Book not found",
	})
}

// DeleteBook handles DELETE /api/books/:id - deletes a book
func (h *APIHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	h.booksMutex.Lock()
	defer h.booksMutex.Unlock()

	for i, book := range h.books {
		if book.ID == id {
			// Remove book from slice
			h.books = append(h.books[:i], h.books[i+1:]...)
			return c.Status(fiber.StatusNoContent).Send(nil)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Book not found",
	})
}
