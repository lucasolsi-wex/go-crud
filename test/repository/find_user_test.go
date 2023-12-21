package repository

import (
	"fmt"
	"github.com/lucasolsi-wex/go-crud/src/model/entity"
	"github.com/lucasolsi-wex/go-crud/src/repository"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestFindUserRepository(t *testing.T) {
	databaseName := "database_test"
	collectionName := "collection_test"
	viper.Set("MONGODB_DATABASE_COLLECTION", collectionName)
	defer viper.Reset()

	mongoTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mongoTestDb.ClearCollections()

	mongoTestDb.Run("find_user_by_id", func(t *mtest.T) {
		userEntity := entity.UserEntity{
			Id:        primitive.ObjectID{},
			Email:     "email@email.com",
			FirstName: "Test",
			LastName:  "Another test",
			Age:       25,
		}
		t.AddMockResponses(mtest.CreateCursorResponse(
			1, fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch, convertEntityToBson(userEntity)))

		mockDb := t.Client.Database(databaseName)

		repo := repository.NewUserRepository(mockDb)
		userDomain, err := repo.FindUserById(userEntity.Id.Hex())
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), userEntity.Id.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetFirstName(), userEntity.FirstName)
		assert.EqualValues(t, userDomain.GetLastName(), userEntity.LastName)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mongoTestDb.Run("user_not_found", func(t *mtest.T) {
		t.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		mockDb := t.Client.Database(databaseName)

		repo := repository.NewUserRepository(mockDb)
		userDomain, err := repo.FindUserById("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.Id},
		{Key: "firstName", Value: userEntity.FirstName},
		{Key: "lastName", Value: userEntity.LastName},
		{Key: "email", Value: userEntity.Email},
		{Key: "age", Value: userEntity.Age},
	}
}
