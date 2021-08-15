package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID      primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}
