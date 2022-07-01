package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Library struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
}
