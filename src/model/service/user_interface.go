package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
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

type UserDomainService interface {
	CreateUser(userModel model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr)
	FindUserById(id string) (*models.UserResponse, *custom_errors.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
