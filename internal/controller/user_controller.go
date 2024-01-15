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

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(gc *gin.Context)
	FindUserById(gc *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func (uc *userControllerInterface) CreateUser(gc *gin.Context) {
	var userRequest models.UserRequest

	if err := gc.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object! Error=%s\n", err.Error())
		customErr := validation.ValidateUserError(err)
		gc.JSON(customErr.Code, customErr)
		return
	}

	customErrName := validation.ValidateFirstAndLastName(userRequest)
	if customErrName != nil {
		gc.JSON(customErrName.Code, customErrName)
		return
	}

	existsNameCombination := uc.service.ExistsByFirstNameAndLastName(userRequest.FirstName, userRequest.LastName)
	customErrUniqueName := validation.ValidateNameUniqueness(existsNameCombination)
	if customErrUniqueName != nil {
		gc.JSON(customErrUniqueName.Code, customErrUniqueName)
		return
	}

	domainResult, err := uc.service.CreateUser(userRequest)
	if err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusOK, domainResult)
}

func (uc *userControllerInterface) FindUserById(gc *gin.Context) {
	idToSearch := gc.Param("userId")

	if _, err := primitive.ObjectIDFromHex(idToSearch); err != nil {
		errorMessage := models.NewUserNotFoundError("Invalid id")
		gc.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserById(idToSearch)
	if err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusOK, userDomain)
}
