package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/converter"
	"github.com/lucasolsi-wex/go-crud/src/model/entity"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (userRepo *userRepository) FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with ID: %s", id)

			return nil, custom_errors.NewUserNotFoundError(errorMessage)
		}
		errorMessage := "Error in Find User By Id"
		return nil, custom_errors.NewInternalServerError(errorMessage)
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}
