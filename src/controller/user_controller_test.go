package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/models"
	"github.com/lucasolsi-wex/go-crud/src/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserControllerInterface_FindUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := service.NewMockUserDomainService(ctrl)
	userController := NewUserControllerInterface(userService)

	t.Run("search user that doesn't exist", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "101010",
			},
		}

		RunRequest(context, param, url.Values{}, "GET", nil)
		userController.FindUserById(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("get user by id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userId := primitive.NewObjectID().Hex()

		param := []gin.Param{{
			Key:   "userId",
			Value: userId,
		}}

		userService.EXPECT().FindUserById(userId).Return(models.UserResponse{
			Id:        "101010",
			FirstName: "First",
			LastName:  "Name",
			Email:     "name@email.com",
			Age:       99,
		}, nil)
		RunRequest(context, param, url.Values{}, "GET", nil)
		userController.FindUserById(context)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	return ctx
}

func RunRequest(gc *gin.Context, params gin.Params, values url.Values, method string, body io.ReadCloser) {
	gc.Request.Method = method
	gc.Request.Header.Set("Content-Type", "application/json")
	gc.Params = params
	gc.Request.Body = body
	gc.Request.URL.RawQuery = values.Encode()
}
