package view

import (
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/models"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) models.UserResponse {
	return models.UserResponse{
		Id:        userDomain.GetId(),
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Age:       userDomain.GetAge(),
	}
}
