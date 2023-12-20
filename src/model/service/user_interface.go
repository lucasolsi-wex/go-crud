package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/repository"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr)
	FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
