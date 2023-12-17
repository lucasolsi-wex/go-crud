package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (uc *userControllerInterface) FindUserById(gc *gin.Context) {
	idToSearch := gc.Param("userId")

	if _, err := primitive.ObjectIDFromHex(idToSearch); err != nil {
		errorMessage := custom_errors.NewBadRequestError("Invalid id")
		gc.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserById(idToSearch)
	if err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
