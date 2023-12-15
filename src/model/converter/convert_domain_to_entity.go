package converter

import (
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:     domain.GetEmail(),
		FirstName: domain.GetFirstName(),
		LastName:  domain.GetLastName(),
		Age:       domain.GetAge(),
	}
}
