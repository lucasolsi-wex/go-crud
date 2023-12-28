package service

import (
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/service"
	"github.com/lucasolsi-wex/go-crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := mocks.NewMockUserRepository(controller)
	userService := service.NewUserDomainService(repository)

	t.Run("if_user_with_id_exists_then_returns_successfully", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("First", "Name", "email@email.com", 34)
		userDomain.SetId(id)

		repository.EXPECT().FindUserById(id).Return(userDomain, nil)

		userDomainReturn, err := userService.FindUserById(id)
		assert.Nil(t, err)
		assert.NotNil(t, userDomainReturn)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetFirstName(), userDomain.GetFirstName())
		assert.EqualValues(t, userDomainReturn.GetLastName(), userDomain.GetLastName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("if_user_does_not_exists_then_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserById(id).Return(nil, custom_errors.NewUserNotFoundError("User not found!"))

		userDomainReturn, err := userService.FindUserById(id)
		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "User not found!")
	})
}
