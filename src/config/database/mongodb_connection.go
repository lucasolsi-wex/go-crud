package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MongoUrl      = "MONGODB_URL"
	MongoDatabase = "MONGODB_DATABASE_NAME"
)

func NewMongoDBConnection(
	ctx context.Context) (*mongo.Database, error) {
	databaseUri := os.Getenv(MongoUrl)
	databaseName := os.Getenv(MongoDatabase)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database(databaseName), nil
}
