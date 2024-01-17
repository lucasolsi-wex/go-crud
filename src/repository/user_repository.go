package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MongoDBUserDb = "MONGODB_DATABASE_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func (userRepo *userRepository) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	count, _ := collection.CountDocuments(context.Background(), bson.M{"firstName": firstName, "lastName": lastName})

	if count >= 1 {
		return true
	}
	return false
}

func (userRepo *userRepository) FindUserById(id string) (models.UserResponse, *custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	userResponse := &models.UserResponse{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userResponse)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with ID: %s", id)

			return models.UserResponse{}, custom_errors.NewUserNotFoundError(errorMessage)
		}
		errorMessage := "Error in Find User By Id"
		return models.UserResponse{}, custom_errors.NewInternalServerError(errorMessage)
	}

	return models.UserResponse{Id: userResponse.Id, FirstName: userResponse.FirstName, LastName: userResponse.LastName,
		Email: userResponse.Email, Age: userResponse.Age}, nil
}

func (userRepo *userRepository) CreateUser(request models.UserRequest) (*models.UserResponse, *custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	entity := models.NewUser(request.FirstName, request.LastName, request.Email, request.Age)

	result, err := collection.InsertOne(context.Background(), entity)

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err.Error())
	}

	entity.Id = result.InsertedID.(primitive.ObjectID)

	return models.FromEntity(entity), nil
}

type UserRepository interface {
	CreateUser(request models.UserRequest) (*models.UserResponse, *custom_errors.CustomErr)
	FindUserById(id string) (models.UserResponse, *custom_errors.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
