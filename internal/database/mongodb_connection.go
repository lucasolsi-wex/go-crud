package database

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoUrl      = "MONGODB_URL"
	MongoDatabase = "MONGODB_DATABASE_NAME"
	MongoUser     = "MONGODB_USERNAME"
	MongoPassword = "MONGODB_PASSWORD"
)

func NewMongoDBConnection(
	ctx context.Context) (*mongo.Database, error) {

	databaseUri := viper.GetString(MongoUrl)
	databaseName := viper.GetString(MongoDatabase)
	mongoUser := viper.GetString(MongoUser)
	mongoPassword := viper.GetString(MongoPassword)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUri).SetAuth(options.Credential{
		AuthMechanism:           "",
		AuthMechanismProperties: nil,
		AuthSource:              "",
		Username:                mongoUser,
		Password:                mongoPassword,
		PasswordSet:             false,
	}))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database(databaseName), nil
}
