package model

type UserDomainInterface interface {
	GetEmail() string
	GetFirstName() string
	GetLastName() string
	GetAge() int8
}

func NewUserDomain(firstName, lastName, email string, age int8) UserDomainInterface {
	return &userDomain{
		firstName,
		lastName,
		email,
		age,
	}
}

type userDomain struct {
	firstName string
	lastName  string
	email     string
	age       int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetFirstName() string {
	return ud.firstName
}
func (ud *userDomain) GetLastName() string {
	return ud.lastName
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
