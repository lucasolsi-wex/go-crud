package repository

import (
	"context"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/converter"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MongoDBUserDb = "MONGODB_DATABASE_COLLECTION"
)

func (userRepo *userRepository) CreateUser(domainInterface model.UserDomainInterface) (
	model.UserDomainInterface, *custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(domainInterface)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, custom_errors.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
