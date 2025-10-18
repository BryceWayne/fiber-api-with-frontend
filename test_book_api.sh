#!/bin/bash

# Book Store API Test Script
# This script demonstrates all CRUD operations for the Book API

BASE_URL="http://localhost:3000"

echo "=========================================="
echo "Book Store API - CRUD Operations Test"
echo "=========================================="
echo ""

echo "1. GET /api/books - List all books (should be empty initially)"
curl -s -X GET "${BASE_URL}/api/books" | python3 -m json.tool
echo ""
echo ""

echo "2. POST /api/books - Create first book"
curl -s -X POST "${BASE_URL}/api/books" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan Donovan and Brian Kernighan",
    "isbn": "978-0134190440",
    "year": 2015
  }' | python3 -m json.tool
echo ""
echo ""

echo "3. POST /api/books - Create second book"
curl -s -X POST "${BASE_URL}/api/books" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "isbn": "978-0132350884",
    "year": 2008
  }' | python3 -m json.tool
echo ""
echo ""

echo "4. POST /api/books - Create third book"
curl -s -X POST "${BASE_URL}/api/books" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Design Patterns",
    "author": "Gang of Four",
    "isbn": "978-0201633610",
    "year": 1994
  }' | python3 -m json.tool
echo ""
echo ""

echo "5. GET /api/books - List all books"
curl -s -X GET "${BASE_URL}/api/books" | python3 -m json.tool
echo ""
echo ""

echo "6. GET /api/books/2 - Get specific book by ID"
curl -s -X GET "${BASE_URL}/api/books/2" | python3 -m json.tool
echo ""
echo ""

echo "7. PUT /api/books/2 - Update book"
curl -s -X PUT "${BASE_URL}/api/books/2" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code: A Handbook of Agile Software Craftsmanship",
    "author": "Robert C. Martin",
    "isbn": "978-0132350884",
    "year": 2008
  }' | python3 -m json.tool
echo ""
echo ""

echo "8. GET /api/books - Verify update"
curl -s -X GET "${BASE_URL}/api/books" | python3 -m json.tool
echo ""
echo ""

echo "9. DELETE /api/books/3 - Delete book"
curl -s -X DELETE "${BASE_URL}/api/books/3" -w "HTTP Status: %{http_code}\n"
echo ""
echo ""

echo "10. GET /api/books - Verify deletion"
curl -s -X GET "${BASE_URL}/api/books" | python3 -m json.tool
echo ""
echo ""

echo "=========================================="
echo "Error Handling Tests"
echo "=========================================="
echo ""

echo "11. GET /api/books/999 - Test 404 Not Found"
curl -s -X GET "${BASE_URL}/api/books/999" -w "\nHTTP Status: %{http_code}\n"
echo ""
echo ""

echo "12. POST /api/books - Test validation (missing title)"
curl -s -X POST "${BASE_URL}/api/books" \
  -H "Content-Type: application/json" \
  -d '{"author": "Test Author"}' \
  -w "\nHTTP Status: %{http_code}\n"
echo ""
echo ""

echo "13. PUT /api/books/abc - Test invalid ID"
curl -s -X PUT "${BASE_URL}/api/books/abc" \
  -H "Content-Type: application/json" \
  -d '{"title": "Test", "author": "Test"}' \
  -w "\nHTTP Status: %{http_code}\n"
echo ""
echo ""

echo "=========================================="
echo "All tests completed!"
echo "=========================================="
