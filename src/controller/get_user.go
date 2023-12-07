package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
)

func GetUserById(gc *gin.Context) {
	err := custom_errors.NewBadRequestError("Wrong!")
	gc.JSON(err.Code, err)
}
