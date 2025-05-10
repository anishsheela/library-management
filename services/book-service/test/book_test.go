package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "book-service/internal/models"
	"book-service/internal/controllers"
    "github.com/gorilla/mux"
)

// Test AddBook
func TestAddBook(t *testing.T) {
    newBook := models.Book{BookID: "123", Title: "Test Book", Author: "Author", Available: true}
    jsonBook, _ := json.Marshal(newBook)

    req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBook))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    controllers.AddBook(rr, req)

    if rr.Code != http.StatusCreated {
        t.Errorf("Expected status Created, got %v", rr.Code)
    }
}

// Test GetBooks
func TestGetBooks(t *testing.T) {
    req, _ := http.NewRequest("GET", "/books", nil)
    rr := httptest.NewRecorder()

    controllers.GetBooks(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status OK, got %v", rr.Code)
    }
}

// Test GetBookByID
func TestGetBookByID(t *testing.T) {
    req, _ := http.NewRequest("GET", "/books/123", nil)
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/books/{bookId}", controllers.GetBookByID)
    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status OK, got %v", rr.Code)
    }
}

// Test UpdateBookAvailability
func TestUpdateBookAvailability(t *testing.T) {
    updateData := map[string]bool{"available": false}
    jsonUpdate, _ := json.Marshal(updateData)

    req, _ := http.NewRequest("PUT", "/books/123", bytes.NewBuffer(jsonUpdate))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    controllers.UpdateBookAvailability(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status OK, got %v", rr.Code)
    }
}

// Test DeleteBook
func TestDeleteBook(t *testing.T) {
    req, _ := http.NewRequest("DELETE", "/books/123", nil)
    rr := httptest.NewRecorder()

    controllers.DeleteBook(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status OK, got %v", rr.Code)
    }
}
