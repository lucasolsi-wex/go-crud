package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/models"
	"github.com/lucasolsi-wex/go-crud/src/repository"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

func (ud *userDomainService) FindUserById(id string) (*models.UserResponse, *custom_errors.CustomErr) {
	return ud.repository.FindUserById(id)
}

func (ud *userDomainService) CreateUser(request models.UserRequest) (*models.UserResponse, *custom_errors.CustomErr) {
	userFromRepo, err := ud.repository.CreateUser(request)

	if err != nil {
		return nil, err
	}

	return userFromRepo, nil
}

type UserDomainService interface {
	CreateUser(request models.UserRequest) (*models.UserResponse, *custom_errors.CustomErr)
	FindUserById(id string) (*models.UserResponse, *custom_errors.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
