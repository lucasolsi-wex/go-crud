package service

import (
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserInterface(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
}
