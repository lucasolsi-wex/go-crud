package service

import (
	"errors"
	"fmt"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/lucasolsi-wex/go-crud/internal/validation"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserInterface(repository repository.UserRepository) UserInterfaceService {
	return &userInterfaceService{repository}
}

type userInterfaceService struct {
	repository repository.UserRepository
}

func (us *userInterfaceService) ExistsByFirstNameAndLastName(firstName, lastName string) (bool, *models.CustomErr) {
	uniquenessCheck := us.repository.ExistsByFirstNameAndLastName(firstName, lastName)
	if uniquenessCheck {
		return false, validation.ValidateNameUniqueness(uniquenessCheck)
	}
	return uniquenessCheck, nil
}

func (us *userInterfaceService) FindUserById(id string) (*models.UserResponse, *models.CustomErr) {
	existingUser, err := us.repository.FindUserById(id)
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

func (us *userInterfaceService) CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr) {
	userToRepo := models.NewUser(request.FirstName, request.LastName, request.Email, request.Age)
	userFromRepo, err := us.repository.CreateUser(userToRepo)

	if err != nil {
		return nil, models.NewInternalServerError(err.Error())
	}

	return models.FromEntity(*userFromRepo), nil
}

type UserInterfaceService interface {
	CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr)
	FindUserById(id string) (*models.UserResponse, *models.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) (bool, *models.CustomErr)
}
