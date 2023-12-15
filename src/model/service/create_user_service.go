package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr) {

	userDomainRepo, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepo, nil
}
