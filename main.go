package main

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Book represents a book in our library
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Year        int       `json:"year"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// BookStore manages our books in memory
type BookStore struct {
	books  map[int]*Book
	nextID int
	mu     sync.RWMutex
}

// NewBookStore creates a new book store
func NewBookStore() *BookStore {
	store := &BookStore{
		books:  make(map[int]*Book),
		nextID: 1,
	}
	// Add some initial books
	store.AddBook(&Book{
		Title:       "The Go Programming Language",
		Author:      "Alan A. A. Donovan & Brian W. Kernighan",
		Year:        2015,
		Description: "A comprehensive guide to Go programming",
	})
	store.AddBook(&Book{
		Title:       "Clean Code",
		Author:      "Robert C. Martin",
		Year:        2008,
		Description: "A handbook of agile software craftsmanship",
	})
	store.AddBook(&Book{
		Title:       "Design Patterns",
		Author:      "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides",
		Year:        1994,
		Description: "Elements of reusable object-oriented software",
	})
	return store
}

// AddBook adds a new book to the store
func (bs *BookStore) AddBook(book *Book) *Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	book.ID = bs.nextID
	book.CreatedAt = time.Now()
	bs.books[book.ID] = book
	bs.nextID++

	return book
}

// GetBook retrieves a book by ID
func (bs *BookStore) GetBook(id int) (*Book, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	book, exists := bs.books[id]
	return book, exists
}

// GetAllBooks returns all books
func (bs *BookStore) GetAllBooks() []*Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	books := make([]*Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}
	return books
}

// UpdateBook updates an existing book
func (bs *BookStore) UpdateBook(id int, updated *Book) (*Book, bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	book, exists := bs.books[id]
	if !exists {
		return nil, false
	}

	book.Title = updated.Title
	book.Author = updated.Author
	book.Year = updated.Year
	book.Description = updated.Description

	return book, true
}

// DeleteBook removes a book from the store
func (bs *BookStore) DeleteBook(id int) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	_, exists := bs.books[id]
	if !exists {
		return false
	}

	delete(bs.books, id)
	return true
}

func main() {
	// Initialize template engine
	engine := html.New("./views", ".html")

	// Create Fiber app with template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initialize book store
	store := NewBookStore()

	// Serve static files
	app.Static("/static", "./static")

	// Routes

	// Home page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Books Library",
		})
	})

	// API Routes

	// Get all books (JSON)
	app.Get("/api/books", func(c *fiber.Ctx) error {
		books := store.GetAllBooks()
		return c.JSON(books)
	})

	// Get all books (HTML fragment for htmx)
	app.Get("/books", func(c *fiber.Ctx) error {
		books := store.GetAllBooks()
		return c.Render("partials/book-list", fiber.Map{
			"Books": books,
		})
	})

	// Get single book
	app.Get("/api/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		book, exists := store.GetBook(id)
		if !exists {
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}

		return c.JSON(book)
	})

	// Get edit form for a book
	app.Get("/books/:id/edit", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid book ID")
		}

		book, exists := store.GetBook(id)
		if !exists {
			return c.Status(404).SendString("Book not found")
		}

		return c.Render("partials/book-edit", fiber.Map{
			"Book": book,
		})
	})

	// Create new book
	app.Post("/api/books", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		newBook := store.AddBook(book)
		return c.Status(201).JSON(newBook)
	})

	// Create new book (from form)
	app.Post("/books", func(c *fiber.Ctx) error {
		year, _ := strconv.Atoi(c.FormValue("year"))

		book := &Book{
			Title:       c.FormValue("title"),
			Author:      c.FormValue("author"),
			Year:        year,
			Description: c.FormValue("description"),
		}

		store.AddBook(book)

		// Return updated book list
		books := store.GetAllBooks()
		return c.Render("partials/book-list", fiber.Map{
			"Books": books,
		})
	})

	// Update book
	app.Put("/api/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		updatedBook, exists := store.UpdateBook(id, book)
		if !exists {
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}

		return c.JSON(updatedBook)
	})

	// Update book (from form)
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid book ID")
		}

		year, _ := strconv.Atoi(c.FormValue("year"))

		book := &Book{
			Title:       c.FormValue("title"),
			Author:      c.FormValue("author"),
			Year:        year,
			Description: c.FormValue("description"),
		}

		_, exists := store.UpdateBook(id, book)
		if !exists {
			return c.Status(404).SendString("Book not found")
		}

		// Return updated book list
		books := store.GetAllBooks()
		return c.Render("partials/book-list", fiber.Map{
			"Books": books,
		})
	})

	// Delete book
	app.Delete("/api/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		if !store.DeleteBook(id) {
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}

		return c.SendStatus(204)
	})

	// Delete book (htmx)
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid book ID")
		}

		if !store.DeleteBook(id) {
			return c.Status(404).SendString("Book not found")
		}

		// Return updated book list
		books := store.GetAllBooks()
		return c.Render("partials/book-list", fiber.Map{
			"Books": books,
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
