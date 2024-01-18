package service

import (
	"errors"
	"fmt"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/lucasolsi-wex/go-crud/internal/validation"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

func (ud *userDomainService) ExistsByFirstNameAndLastName(firstName, lastName string) (bool, *models.CustomErr) {
	uniquenessCheck := ud.repository.ExistsByFirstNameAndLastName(firstName, lastName)
	if uniquenessCheck {
		return false, validation.ValidateNameUniqueness(uniquenessCheck)
	}
	return uniquenessCheck, nil
}

func (ud *userDomainService) FindUserById(id string) (*models.UserResponse, *models.CustomErr) {
	existingUser, err := ud.repository.FindUserById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with ID: %s", id)

			return nil, models.NewUserNotFoundError(errorMessage)
		}
		errorMessage := "Error in Find User By Id"
		return nil, models.NewInternalServerError(errorMessage)
	}

	return models.FromEntity(*existingUser), nil
}

func (ud *userDomainService) CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr) {
	userToRepo := models.NewUser(request.FirstName, request.LastName, request.Email, request.Age)
	userFromRepo, err := ud.repository.CreateUser(userToRepo)

	if err != nil {
		return nil, models.NewInternalServerError(err.Error())
	}

	return models.FromEntity(*userFromRepo), nil
}

type UserDomainService interface {
	CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr)
	FindUserById(id string) (*models.UserResponse, *models.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) (bool, *models.CustomErr)
}
