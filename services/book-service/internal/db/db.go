package db

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var BookCollection *mongo.Collection

func InitDB() {
    clientOptions := options.Client().ApplyURI("mongodb://book-db:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    BookCollection = client.Database("book_service").Collection("books")
}
