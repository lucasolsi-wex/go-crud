package response

type UserResponse struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int8   `json:"age,omitempty"`
}
