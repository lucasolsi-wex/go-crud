package request

type UserRequest struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty" binding:"required,email"`
	Age       int8   `json:"age,omitempty" binding:"required,min=18,max=120"`
}
