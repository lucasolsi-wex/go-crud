package models

type UserResponse struct {
	Id        string `json:"id,omitempty" bson:"_id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int8   `json:"age,omitempty"`
}

func FromEntity(entity UserModel) *UserResponse {
	return &UserResponse{
		Id:        entity.Id.Hex(),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Age:       entity.Age,
	}
}
