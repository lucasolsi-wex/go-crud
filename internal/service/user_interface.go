package service

import (
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

func (ud *userDomainService) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	return ud.repository.ExistsByFirstNameAndLastName(firstName, lastName)
}

func (ud *userDomainService) FindUserById(id string) (models.UserResponse, *models.CustomErr) {
	return ud.repository.FindUserById(id)
}

func (ud *userDomainService) CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr) {
	userFromRepo, err := ud.repository.CreateUser(request)

	if err != nil {
		return nil, err
	}

	return userFromRepo, nil
}

type UserDomainService interface {
	CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr)
	FindUserById(id string) (models.UserResponse, *models.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
