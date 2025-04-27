package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "borrowing-service/internal/controllers"
)

func main() {
    // Initialize the router
    router := mux.NewRouter()

    // Register routes (to be implemented later)
    router.HandleFunc("/borrow", controllers.BorrowBook).Methods("POST")
    router.HandleFunc("/return", controllers.ReturnBook).Methods("POST")
    router.HandleFunc("/borrowings/{userId}", controllers.GetBorrowingsByUser).Methods("GET")
    router.HandleFunc("/borrowings/overdue", controllers.GetOverdueBorrowings).Methods("GET")

    fmt.Println("Borrowing Service running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

