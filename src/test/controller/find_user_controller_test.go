package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/controller"
	"github.com/lucasolsi-wex/go-crud/src/model"
	"github.com/lucasolsi-wex/go-crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserControllerInterface_FindUserById(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	service := mocks.NewMockUserDomainService(mockController)
	userController := controller.NewUserControllerInterface(service)

	t.Run("if_user_id_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "101010",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		userController.FindUserById(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("if user id is valid returns user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserById(id).Return(
			model.NewUserDomain("First", "Name", "test@test.com", 32), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
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

func MakeRequest(
	gc *gin.Context, param gin.Params, u url.Values, method string, body io.ReadCloser) {
	gc.Request.Method = method
	gc.Request.Header.Set("Content-Type", "application/json")
	gc.Params = param

	gc.Request.Body = body
	gc.Request.URL.RawQuery = u.Encode()
}
