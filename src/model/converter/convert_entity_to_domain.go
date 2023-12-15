package converter

import (
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.FirstName, entity.LastName, entity.Email, entity.Age)
	domain.SetId(entity.Id.Hex())

	return domain
}
