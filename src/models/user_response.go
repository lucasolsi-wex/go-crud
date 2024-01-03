package models

type UserResponse struct {
	Id        string `json:"id,omitempty" bson:"_id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int8   `json:"age,omitempty"`
}
