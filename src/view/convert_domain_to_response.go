package view

import (
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/response"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:        "",
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Age:       userDomain.GetAge(),
	}
}
