package repository

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr)
	FindUserById(id string) (*models.UserResponse, *custom_errors.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
