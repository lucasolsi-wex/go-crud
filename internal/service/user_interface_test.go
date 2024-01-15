package service

import (
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserInterface(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repoMock := repository.NewMockUserRepository(controller)
	userService := NewUserDomainService(repoMock)

	t.Run("if_first_and_last_name_already_exists_returns_true", func(t *testing.T) {
		repoMock.EXPECT().ExistsByFirstNameAndLastName("first", "name").Return(true)

		exists := userService.ExistsByFirstNameAndLastName("first", "name")
		assert.EqualValues(t, exists, true)
	})
}
