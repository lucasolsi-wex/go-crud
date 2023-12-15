package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	FirstName string             `bson:"firstName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
	Age       int8               `bson:"age,omitempty"`
}
