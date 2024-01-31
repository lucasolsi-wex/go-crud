package models

type UserRequest struct {
	FirstName string `json:"firstName,omitempty" binding:"required,alpha"`
	LastName  string `json:"lastName,omitempty" binding:"required,alpha"`
	Email     string `json:"email,omitempty" binding:"required,email"`
	Age       int8   `json:"age,omitempty" binding:"required,min=18,max=120"`
}
