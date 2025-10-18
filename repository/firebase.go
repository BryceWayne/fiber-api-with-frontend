package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// FirebaseRepository implements BookRepository using Google Firestore
type FirebaseRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirebaseRepository creates a new Firebase repository
// projectID is your Google Cloud project ID
// collection is the Firestore collection name (e.g., "books")
func NewFirebaseRepository(ctx context.Context, projectID, collection string) (*FirebaseRepository, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create Firestore client: %w", err)
	}
	
	return &FirebaseRepository{
		client:     client,
		collection: collection,
	}, nil
}

// Close closes the Firestore client connection
func (r *FirebaseRepository) Close() error {
	return r.client.Close()
}

// Create adds a new book to the Firestore repository
func (r *FirebaseRepository) Create(ctx context.Context, book *Book) error {
	// Generate a new document reference with auto-generated ID
	docRef := r.client.Collection(r.collection).NewDoc()
	book.ID = docRef.ID
	
	_, err := docRef.Set(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}
	
	return nil
}

// GetAll retrieves all books from the Firestore repository
func (r *FirebaseRepository) GetAll(ctx context.Context) ([]Book, error) {
	iter := r.client.Collection(r.collection).Documents(ctx)
	defer iter.Stop()
	
	var books []Book
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate books: %w", err)
		}
		
		var book Book
		if err := doc.DataTo(&book); err != nil {
			return nil, fmt.Errorf("failed to parse book data: %w", err)
		}
		
		// Ensure ID is set from document reference
		book.ID = doc.Ref.ID
		books = append(books, book)
	}
	
	return books, nil
}

// GetByID retrieves a specific book by its ID from Firestore
func (r *FirebaseRepository) GetByID(ctx context.Context, id string) (*Book, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}
	
	var book Book
	if err := doc.DataTo(&book); err != nil {
		return nil, fmt.Errorf("failed to parse book data: %w", err)
	}
	
	// Ensure ID is set from document reference
	book.ID = doc.Ref.ID
	return &book, nil
}

// Update modifies an existing book in the Firestore repository
func (r *FirebaseRepository) Update(ctx context.Context, id string, book *Book) error {
	book.ID = id
	
	// Use Update() which fails if document doesn't exist, avoiding extra read
	_, err := r.client.Collection(r.collection).Doc(id).Set(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}
	
	return nil
}

// Delete removes a book from the Firestore repository
func (r *FirebaseRepository) Delete(ctx context.Context, id string) error {
	// Delete the document - Firestore handles non-existent documents gracefully
	_, err := r.client.Collection(r.collection).Doc(id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}
	
	return nil
}
