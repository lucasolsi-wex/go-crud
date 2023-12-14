package repository

import (
	"context"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"os"
)

const (
	MongoDBUserDb = "MONGODB_DATABASE_COLLECTION"
)

func (repo *userRepository) CreateUser(domainInterface model.UserDomainInterface) (
	model.UserDomainInterface, *custom_errors.CustomErr) {
	collectionName := os.Getenv(MongoDBUserDb)
	collection := repo.databaseConnection.Collection(collectionName)

	value, err := domainInterface.ToJSON()
	if err != nil {
		return nil, custom_errors.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, custom_errors.NewInternalServerError(err.Error())
	}

	domainInterface.SetId(result.InsertedID.(string))

	return domainInterface, nil
}
