package controllers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"github.com/gorilla/mux"
)

type BorrowingInfo struct {
	BookID    string    `json:"bookId"`
	DueDate   time.Time `json:"dueDate"`
}

var Borrowings = make(map[string][]BorrowingInfo)
var mutex = &sync.Mutex{}

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID string `json:"userId"`
		BookID string `json:"bookId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	for _, b := range Borrowings[request.UserID] {
		if b.BookID == request.BookID {
			http.Error(w, "Book already borrowed by user", http.StatusConflict)
			return
		}
	}
	dueDate := time.Now().Add(14 * 24 * time.Hour)
	Borrowings[request.UserID] = append(Borrowings[request.UserID], BorrowingInfo{
		BookID:    request.BookID,
		DueDate:   dueDate,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book borrowed successfully"})
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID string `json:"userId"`
		BookID string `json:"bookId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	userBorrowings := Borrowings[request.UserID]
	for i, b := range userBorrowings {
		if b.BookID == request.BookID {
			Borrowings[request.UserID] = append(userBorrowings[:i], userBorrowings[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Book returned successfully"})
			return
		}
	}
	http.Error(w, "Book not found for the user", http.StatusNotFound)
}

func GetBorrowingsByUser(w http.ResponseWriter, r *http.Request) {
    userID := mux.Vars(r)["userId"]
    mutex.Lock()
    defer mutex.Unlock()
    userBorrowings, exists := Borrowings[userID]
    if !exists || len(userBorrowings) == 0 {
        http.Error(w, "No borrowings found for the user", http.StatusNotFound)
        return
    }
    type ResponseBorrowingInfo struct {
        BookID    string    `json:"bookId"`
        DueDate   time.Time `json:"dueDate"`
        IsOverdue bool      `json:"isOverdue"`
    }
    var response []ResponseBorrowingInfo
    for _, b := range userBorrowings {
        response = append(response, ResponseBorrowingInfo{
            BookID:    b.BookID,
            DueDate:   b.DueDate,
            IsOverdue: b.DueDate.Before(time.Now()),
        })
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func GetOverdueBorrowings(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()
    overdueBorrowings := make(map[string][]BorrowingInfo) // Map grouped by userId
    currentTime := time.Now()
    for userID, userBorrowings := range Borrowings {
        for _, b := range userBorrowings {
            if b.DueDate.Before(currentTime) {
                overdueBorrowings[userID] = append(overdueBorrowings[userID], b)
            }
        }
    }
    if len(overdueBorrowings) == 0 {
        http.Error(w, "No overdue borrowings found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(overdueBorrowings)
}


