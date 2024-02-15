package service

import (
	"context"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := models.UserRequest{
		FirstName: "first",
		LastName:  "name",
		Email:     "email@email.com",
		Age:       33,
	}

	userToRepo := models.UserModel{
		Id:        primitive.ObjectID{},
		Email:     "email@email.com",
		FirstName: "first",
		LastName:  "name",
		Age:       33,
	}

	mockRepo := repository.NewMockUserRepository(ctrl)
	testService := UserService{Repository: mockRepo}

	t.Run("if first and last name already exists then return error", func(t *testing.T) {
		mockRepo.EXPECT().ExistsByFirstNameAndLastName("first", "name", context.Background()).Return(true, nil)

		result, err := testService.CreateUser(mockUser, context.Background())
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})

	t.Run("if doesnt exists first and last name then return successful user creation", func(t *testing.T) {

		mockRepo.EXPECT().ExistsByFirstNameAndLastName("first", "name", context.Background()).Return(false, nil)
		mockRepo.EXPECT().CreateUser(userToRepo, context.Background()).Return(&userToRepo, nil)

		result, err := testService.CreateUser(mockUser, context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, result)

		assert.NotNil(t, result.Id)
		assert.EqualValues(t, mockUser.FirstName, result.FirstName)
		assert.EqualValues(t, mockUser.LastName, result.LastName)
		assert.EqualValues(t, mockUser.Email, result.Email)
		assert.EqualValues(t, mockUser.Age, result.Age)
	})

	t.Run("if id is valid then return existing user", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("65b7b286736098b80b440c30")
		mockRepo.EXPECT().FindUserById(objectId, context.Background()).Return(&userToRepo, nil)

		result, err := testService.FindUserById("65b7b286736098b80b440c30", context.Background())
		assert.Error(t, err)
		assert.NotNil(t, result)
	})

	t.Run("if there's no matching userId then return error", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("65b7b286736098b80b440c30")
		mockRepo.EXPECT().FindUserById(objectId, context.Background()).Return(nil, mongo.ErrNoDocuments)

		result, err := testService.FindUserById("65b7b286736098b80b440c30", context.Background())
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
