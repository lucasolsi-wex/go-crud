package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/service"
	"github.com/lucasolsi-wex/go-crud/internal/validation"
	"log"
	"net/http"
)

func NewUserControllerInterface(serviceInterface service.UserService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(gc *gin.Context)
	FindUserById(gc *gin.Context)
}

type userControllerInterface struct {
	service service.UserService
}

func (uc *userControllerInterface) CreateUser(gc *gin.Context) {
	var userRequest models.UserRequest

	if err := gc.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object: %v", err.Error())
		customErr := validation.ValidateUserError(err)
		gc.JSON(http.StatusBadRequest, customErr)
		return
	}

	domainResult, err := uc.service.CreateUser(userRequest, gc)
	if err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusCreated, domainResult)
}

func (uc *userControllerInterface) FindUserById(gc *gin.Context) {
	idToSearch := gc.Param("userId")

	userDomain, err := uc.service.FindUserById(idToSearch, gc)
	if err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusOK, userDomain)
}
