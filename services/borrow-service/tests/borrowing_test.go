package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "borrow-service/internal/controllers"
)

func TestBorrowBook(t *testing.T) {
    requestBody := map[string]string{"userId": "user1", "bookId": "book1"}
    body, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/borrow", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    http.HandlerFunc(controllers.BorrowBook).ServeHTTP(rr, req)
    if rr.Code != http.StatusOK {
        t.Errorf("Got status %v, want %v", rr.Code, http.StatusOK)
    }
    expected := `{"message":"Book borrowed successfully"}`
    actual := strings.TrimSpace(rr.Body.String())
    if actual != expected {
        t.Errorf("Got body %v, want %v", actual, expected)
    }
}

func TestReturnBook(t *testing.T) {
    controllers.BorrowBook(httptest.NewRecorder(), createMockRequest("user1", "book1"))
    requestBody := map[string]string{"userId": "user1", "bookId": "book1"}
    body, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/return", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    http.HandlerFunc(controllers.ReturnBook).ServeHTTP(rr, req)
    if rr.Code != http.StatusOK {
        t.Errorf("Got status %v, want %v", rr.Code, http.StatusOK)
    }
    expected := `{"message":"Book returned successfully"}`
    actual := strings.TrimSpace(rr.Body.String())
    if actual != expected {
        t.Errorf("Got body %v, want %v", actual, expected)
    }
}

func createMockRequest(userID, bookID string) *http.Request {
    requestBody := map[string]string{"userId": userID, "bookId": bookID}
    body, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/borrow", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    return req
}
