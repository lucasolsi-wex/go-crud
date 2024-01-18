package service

import (
	"errors"
	"fmt"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/lucasolsi-wex/go-crud/internal/validation"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInterfaceService struct {
	Repository repository.UserRepository
}

func (us UserInterfaceService) ExistsByFirstNameAndLastName(firstName, lastName string) (bool, *models.CustomErr) {
	uniquenessCheck := us.Repository.ExistsByFirstNameAndLastName(firstName, lastName)
	if uniquenessCheck {
		return false, validation.ValidateNameUniqueness(uniquenessCheck)
	}
	return uniquenessCheck, nil
}

func (us UserInterfaceService) FindUserById(id string) (*models.UserResponse, *models.CustomErr) {
	existingUser, err := us.Repository.FindUserById(id)
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

func (us UserInterfaceService) CreateUser(request models.UserRequest) (*models.UserResponse, *models.CustomErr) {
	userToRepo := models.NewUser(request.FirstName, request.LastName, request.Email, request.Age)
	userFromRepo, err := us.Repository.CreateUser(userToRepo)

	if err != nil {
		return nil, models.NewInternalServerError(err.Error())
	}

	return models.FromEntity(*userFromRepo), nil
}
