package repository

import (
	"context"
	"fmt"
	"strconv"
	"sync"
)

// MemoryRepository implements BookRepository using in-memory storage
type MemoryRepository struct {
	books      map[string]Book
	booksMutex sync.RWMutex
	nextID     int
}

// NewMemoryRepository creates a new in-memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		books:  make(map[string]Book),
		nextID: 1,
	}
}

// Create adds a new book to the in-memory repository
func (r *MemoryRepository) Create(ctx context.Context, book *Book) error {
	r.booksMutex.Lock()
	defer r.booksMutex.Unlock()
	
	// Generate ID
	book.ID = strconv.Itoa(r.nextID)
	r.nextID++
	
	r.books[book.ID] = *book
	return nil
}

// GetAll retrieves all books from the in-memory repository
func (r *MemoryRepository) GetAll(ctx context.Context) ([]Book, error) {
	r.booksMutex.RLock()
	defer r.booksMutex.RUnlock()
	
	books := make([]Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	
	return books, nil
}

// GetByID retrieves a specific book by its ID
func (r *MemoryRepository) GetByID(ctx context.Context, id string) (*Book, error) {
	r.booksMutex.RLock()
	defer r.booksMutex.RUnlock()
	
	book, exists := r.books[id]
	if !exists {
		return nil, fmt.Errorf("book not found")
	}
	
	return &book, nil
}

// Update modifies an existing book in the in-memory repository
func (r *MemoryRepository) Update(ctx context.Context, id string, book *Book) error {
	r.booksMutex.Lock()
	defer r.booksMutex.Unlock()
	
	if _, exists := r.books[id]; !exists {
		return fmt.Errorf("book not found")
	}
	
	book.ID = id
	r.books[id] = *book
	return nil
}

// Delete removes a book from the in-memory repository
func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	r.booksMutex.Lock()
	defer r.booksMutex.Unlock()
	
	if _, exists := r.books[id]; !exists {
		return fmt.Errorf("book not found")
	}
	
	delete(r.books, id)
	return nil
}

// Close closes the repository - no-op for memory repository
func (r *MemoryRepository) Close() error {
	return nil
}
