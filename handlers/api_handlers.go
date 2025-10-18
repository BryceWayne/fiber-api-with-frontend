package handlers

import (
	"github.com/BryceWayne/fiber-api-with-frontend/repository"
	"github.com/gofiber/fiber/v2"
)

// APIHandler handles API endpoints
type APIHandler struct {
	repo repository.BookRepository
}

// NewAPIHandler creates a new APIHandler with the given repository
func NewAPIHandler(repo repository.BookRepository) *APIHandler {
	return &APIHandler{
		repo: repo,
	}
}

// Hello returns a hello message
// @Summary Get hello message
// @Description Returns a simple hello message in JSON format
// @Tags api
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Hello message"
// @Router /api/hello [get]
func (h *APIHandler) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello from Fiber API!",
	})
}

// HTML response template for HTMX
const helloHTMLTemplate = `<p class="api-result">API Response: Hello from Fiber API!</p>`

// HelloHTML returns an HTML fragment for HTMX
// @Summary Get hello HTML fragment
// @Description Returns an HTML fragment for HTMX integration
// @Tags api
// @Accept json
// @Produce html
// @Success 200 {string} string "HTML fragment"
// @Router /api/hello-html [get]
func (h *APIHandler) HelloHTML(c *fiber.Ctx) error {
	return c.SendString(helloHTMLTemplate)
}

// CreateBook handles POST /api/books - creates a new book
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body repository.Book true "Book object to create"
// @Success 201 {object} repository.Book "Created book"
// @Failure 400 {object} map[string]string "Invalid request body or missing required fields"
// @Router /api/books [post]
func (h *APIHandler) CreateBook(c *fiber.Ctx) error {
	book := new(repository.Book)

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

	if err := h.repo.Create(c.Context(), book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// GetBooks handles GET /api/books - returns all books
// @Summary Get all books
// @Description Returns a list of all books in the store
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} repository.Book "List of books"
// @Router /api/books [get]
func (h *APIHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.repo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve books",
		})
	}

	return c.JSON(books)
}

// GetBook handles GET /api/books/:id - returns a specific book
// @Summary Get a book by ID
// @Description Returns a specific book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} repository.Book "Book details"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /api/books/{id} [get]
func (h *APIHandler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.JSON(book)
}

// UpdateBook handles PUT /api/books/:id - updates an existing book
// @Summary Update a book
// @Description Updates an existing book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body repository.Book true "Updated book object"
// @Success 200 {object} repository.Book "Updated book"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /api/books/{id} [put]
func (h *APIHandler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedBook := new(repository.Book)
	if err := c.BodyParser(updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if updatedBook.Title == "" || updatedBook.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and Author are required fields",
		})
	}

	if err := h.repo.Update(c.Context(), id, updatedBook); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.JSON(updatedBook)
}

// DeleteBook handles DELETE /api/books/:id - deletes a book
// @Summary Delete a book
// @Description Deletes a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 "Book deleted successfully"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /api/books/{id} [delete]
func (h *APIHandler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
