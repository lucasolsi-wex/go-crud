package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Age       int8               `json:"age"`
}

func NewUser(firstName, lastName, email string, age int8) UserModel {
	return UserModel{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
