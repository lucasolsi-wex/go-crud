package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model/request"
)

func CreateUser(gc *gin.Context) {
	var userRequest request.UserRequest

	if err := gc.ShouldBindJSON(&userRequest); err != nil {
		customErr := custom_errors.NewBadRequestError(
			fmt.Sprintf("Trouble in paradise with fields: %s\n", err.Error()))
		gc.JSON(customErr.Code, customErr)
		return
	}
	fmt.Println(userRequest)
}
