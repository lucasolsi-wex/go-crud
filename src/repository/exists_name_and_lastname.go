package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

func (userRepo *userRepository) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	count, _ := collection.CountDocuments(context.Background(), bson.M{"firstName": firstName, "lastName": lastName})

	if count >= 1 {
		return true
	}
	return false
}
