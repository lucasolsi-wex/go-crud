package service

import (
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/service"
	"github.com/lucasolsi-wex/go-crud/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_CreateUserService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := mocks.NewMockUserRepository(controller)
	userService := service.NewUserDomainService(repository)

	t.Run("if_first_and_last_name_already_exists_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("First", "Name", "email@email.com", 77)
		userDomain.SetId(id)

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := userService.CreateUser(userDomain)
		assert.Nil(t, err)
		assert.NotNil(t, user)
	})
}
