package repository

import (
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/repository"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateUserRepository(t *testing.T) {
	databaseName := "database_test"
	collectionName := "collection_test"
	viper.Set("MONGODB_DATABASE_COLLECTION", collectionName)
	defer viper.Reset()

	mongoTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mongoTestDb.ClearCollections()

	mongoTestDb.Run("success", func(t *mtest.T) {
		t.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := t.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUser(model.NewUserDomain("teste",
			"teste2", "email@email.com", 20))

		_, errId := primitive.ObjectIDFromHex(userDomain.GetId())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "email@email.com")
	})

	mongoTestDb.Run("error_form_database", func(t *mtest.T) {
		t.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := t.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		domain := model.NewUserDomain("teste",
			"teste2", "email@email.com", 20)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
