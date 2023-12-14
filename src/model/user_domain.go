package model

import "encoding/json"

type UserDomainInterface interface {
	GetEmail() string
	GetFirstName() string
	GetLastName() string
	GetAge() int8
	ToJSON() (string, error)
	SetId(string)
}

func NewUserDomain(firstName, lastName, email string, age int8) UserDomainInterface {
	return &userDomain{"",
		firstName,
		lastName,
		email,
		age,
	}
}

type userDomain struct {
	Id        string
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int8   `json:"age,omitempty"`
}

func (ud *userDomain) ToJSON() (string, error) {
	output, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetFirstName() string {
	return ud.FirstName
}
func (ud *userDomain) GetLastName() string {
	return ud.LastName
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) SetId(id string) {
	ud.Id = id
}
