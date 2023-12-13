package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(domainInterface model.UserDomainInterface) *custom_errors.CustomErr
	FindUser(string) (*model.UserDomainInterface, *custom_errors.CustomErr)
}
