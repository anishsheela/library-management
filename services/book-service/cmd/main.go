package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "book-service/internal/controllers"
    "book-service/internal/db"
)

func main() {
    db.InitDB()
    router := mux.NewRouter()

    // Register book service routes
    router.HandleFunc("/books", controllers.AddBook).Methods("POST")                
    router.HandleFunc("/books", controllers.GetBooks).Methods("GET")                
    router.HandleFunc("/books/{bookId}", controllers.GetBookByID).Methods("GET")   
    router.HandleFunc("/books/{bookId}", controllers.UpdateBookAvailability).Methods("PUT") 
    router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
    router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Book Service is healthy")
    }).Methods("GET")

    fmt.Println("Book Service running on port 5001...")
    log.Fatal(http.ListenAndServe(":5001", router))
}
