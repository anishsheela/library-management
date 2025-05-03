package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    BookID    string             `bson:"bookId" json:"bookId"`
    Title     string             `bson:"title" json:"title"`
    Author    string             `bson:"author" json:"author"`
    Available bool               `bson:"available" json:"available"`
}
