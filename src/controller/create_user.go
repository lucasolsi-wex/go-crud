package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/validation"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/model/request"
	"github.com/lucasolsi-wex/go-crud/src/view"
	"log"
	"net/http"
)

func (uc *userControllerInterface) CreateUser(gc *gin.Context) {
	var userRequest request.UserRequest

	if err := gc.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object! Error=%s\n", err.Error())
		customErr := validation.ValidateUserError(err)
		gc.JSON(customErr.Code, customErr)
		return
	}
	fmt.Println(userRequest)

	domain := model.NewUserDomain(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))

}
