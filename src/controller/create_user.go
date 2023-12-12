package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/validation"
	"github.com/lucasolsi-wex/go-crud/src/model/request"
	"log"
)

func CreateUser(gc *gin.Context) {
	var userRequest request.UserRequest

	if err := gc.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object! Error=%s\n", err.Error())
		customErr := validation.ValidateUserError(err)
		gc.JSON(customErr.Code, customErr)
		return
	}
	fmt.Println(userRequest)
}
