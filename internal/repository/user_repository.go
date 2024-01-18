package repository

import (
	"context"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	MongoDBUserDb = "users"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func (userRepo *userRepository) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	collection := userRepo.databaseConnection.Collection(MongoDBUserDb)

	found, _ := collection.Find(context.Background(), bson.M{"firstName": firstName, "lastName": lastName})
	defer func(found *mongo.Cursor, ctx context.Context) {
		err := found.Close(ctx)
		if err != nil {

		}
	}(found, context.Background())

	var results []bson.M
	for found.Next(context.Background()) {
		var result bson.M
		err := found.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)

		if len(results) == 1 {
			return true
		}
	}
	return false
}

func (userRepo *userRepository) FindUserById(id string) (*models.UserModel, error) {
	collection := userRepo.databaseConnection.Collection(MongoDBUserDb)

	existingUser := &models.UserModel{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(existingUser)

	return existingUser, err
}

func (userRepo *userRepository) CreateUser(request models.UserModel) (*models.UserModel, error) {
	collection := userRepo.databaseConnection.Collection(MongoDBUserDb)

	entity := models.NewUser(request.FirstName, request.LastName, request.Email, request.Age)

	result, err := collection.InsertOne(context.Background(), entity)

	entity.Id = result.InsertedID.(primitive.ObjectID)

	return &entity, err
}

type UserRepository interface {
	CreateUser(request models.UserModel) (*models.UserModel, error)
	FindUserById(id string) (*models.UserModel, error)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
