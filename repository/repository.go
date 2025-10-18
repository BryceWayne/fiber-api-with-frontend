package repository

import (
	"context"
)

// Book represents a book in the bookstore
type Book struct {
	ID     string `json:"id" firestore:"id,omitempty"`
	Title  string `json:"title" firestore:"title"`
	Author string `json:"author" firestore:"author"`
	ISBN   string `json:"isbn" firestore:"isbn"`
	Year   int    `json:"year" firestore:"year"`
}

// BookRepository defines the interface for book storage operations
type BookRepository interface {
	// Create adds a new book to the repository
	Create(ctx context.Context, book *Book) error
	
	// GetAll retrieves all books from the repository
	GetAll(ctx context.Context) ([]Book, error)
	
	// GetByID retrieves a specific book by its ID
	GetByID(ctx context.Context, id string) (*Book, error)
	
	// Update modifies an existing book in the repository
	Update(ctx context.Context, id string, book *Book) error
	
	// Delete removes a book from the repository
	Delete(ctx context.Context, id string) error
	
	// Close closes any open connections and cleans up resources
	Close() error
}
