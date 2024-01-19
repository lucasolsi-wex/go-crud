package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/service"
	"github.com/lucasolsi-wex/go-crud/internal/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func NewUserControllerInterface(serviceInterface service.UserInterfaceService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(gc *gin.Context)
	FindUserById(gc *gin.Context)
}

type userControllerInterface struct {
	service service.UserInterfaceService
}

func (uc *userControllerInterface) CreateUser(gc *gin.Context) {
	var userRequest models.UserRequest

	if err := gc.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object! Error=%s\n", err.Error())
		customErr := validation.ValidateUserError(err)
		gc.JSON(http.StatusBadRequest, customErr)
		return
	}

	customErrName := validation.ValidateFirstAndLastName(userRequest)
	if customErrName != nil {
		log.Print("Error while creating user. There's already a match for name and last name",
			customErrName.Error())
		gc.JSON(http.StatusBadRequest, customErrName)
		return
	}

	domainResult, err := uc.service.CreateUser(userRequest)
	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusCreated, domainResult)
}

func (uc *userControllerInterface) FindUserById(gc *gin.Context) {
	idToSearch := gc.Param("userId")

	if _, err := primitive.ObjectIDFromHex(idToSearch); err != nil {
		errorMessage := models.NewBadRequestError("Invalid id")
		gc.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserById(idToSearch)
	if err != nil {
		gc.JSON(http.StatusNotFound, err)
		return
	}

	gc.JSON(http.StatusOK, userDomain)
}
