package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "borrow-service/internal/controllers"
    "borrow-service/internal/db"
)

func main() {
    db.InitDB()
    router := mux.NewRouter()

    router.HandleFunc("/borrow", controllers.BorrowBook).Methods("POST")
    router.HandleFunc("/return", controllers.ReturnBook).Methods("POST")
    router.HandleFunc("/borrowings/{userId}", controllers.GetBorrowingsByUser).Methods("GET")
    router.HandleFunc("/borrowings/overdue", controllers.GetOverdueBorrowings).Methods("GET")

    fmt.Println("Borrow Service running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

