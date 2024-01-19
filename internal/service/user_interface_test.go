package service

import (
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
	testService := UserInterfaceService{Repository: mockRepo}

	t.Run("if first and last name already exists then return error", func(t *testing.T) {
		mockRepo.EXPECT().ExistsByFirstNameAndLastName("first", "name").Return(true)

		result, err := testService.CreateUser(mockUser)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})

	t.Run("if doesnt exists first and last name then return successful user creation", func(t *testing.T) {

		mockRepo.EXPECT().ExistsByFirstNameAndLastName("first", "name").Return(false)
		mockRepo.EXPECT().CreateUser(userToRepo).Return(&userToRepo, nil)

		result, err := testService.CreateUser(mockUser)

		assert.Nil(t, err)
		assert.NotNil(t, result)

		assert.NotNil(t, result.Id)
		assert.EqualValues(t, mockUser.FirstName, result.FirstName)
		assert.EqualValues(t, mockUser.LastName, result.LastName)
		assert.EqualValues(t, mockUser.Email, result.Email)
		assert.EqualValues(t, mockUser.Age, result.Age)
	})

	t.Run("if id is valid then return existing user", func(t *testing.T) {
		mockRepo.EXPECT().FindUserById("1010").Return(&userToRepo, nil)

		result, err := testService.FindUserById("1010")
		assert.Error(t, err)
		assert.NotNil(t, result)
	})

	t.Run("if there's no matching userId then return error", func(t *testing.T) {
		mockRepo.EXPECT().FindUserById("1010").Return(nil, mongo.ErrNoDocuments)

		result, err := testService.FindUserById("1010")
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
