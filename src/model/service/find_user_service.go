package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
)

func (ud *userDomainService) FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr) {
	return ud.repository.FindUserById(id)
}
