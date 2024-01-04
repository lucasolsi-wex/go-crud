package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Email     string             `json:"email,omitempty"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
	Age       int8               `json:"age,omitempty"`
}

func NewUser(firstName, lastName, email string, age int8) UserModel {
	return UserModel{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
