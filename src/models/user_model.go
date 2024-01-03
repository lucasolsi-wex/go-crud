package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Email     string             `json:"email,omitempty"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
	Age       int8               `json:"age,omitempty"`
}
