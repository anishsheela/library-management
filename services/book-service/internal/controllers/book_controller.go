package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "book-service/internal/db"
    "book-service/internal/models"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request format", http.StatusBadRequest)
        return
    }
    book.Available = true
    _, err := db.BookCollection.InsertOne(context.Background(), book)
    if err != nil {
        http.Error(w, "Failed to add book", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Book added successfully"})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    filter := bson.M{}
    if r.URL.Query().Get("availability") == "true" {
        filter = bson.M{"available": true}
    }
    cursor, err := db.BookCollection.Find(context.Background(), filter)
    if err != nil {
        http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())
    var books []models.Book
    if err := cursor.All(context.Background(), &books); err != nil {
        http.Error(w, "Data error", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
    bookID := mux.Vars(r)["bookId"]
    var book models.Book
    err := db.BookCollection.FindOne(context.Background(), bson.M{"bookId": bookID}).Decode(&book)
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func UpdateBookAvailability(w http.ResponseWriter, r *http.Request) {
    bookID := mux.Vars(r)["bookId"]
    var updateData struct {
        Available bool `json:"available"`
    }
    if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
        http.Error(w, "Invalid request format", http.StatusBadRequest)
        return
    }
    _, err := db.BookCollection.UpdateOne(context.Background(), bson.M{"bookId": bookID}, bson.M{"$set": bson.M{"available": updateData.Available}})
    if err != nil {
        http.Error(w, "Failed to update book availability", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Book availability updated successfully"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    bookID := mux.Vars(r)["bookId"]
    _, err := db.BookCollection.DeleteOne(context.Background(), bson.M{"bookId": bookID})
    if err != nil {
        http.Error(w, "Failed to delete book", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Book removed successfully"})
}
