package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/controller"
)

func InitRoutes(rg *gin.RouterGroup, userController controller.UserControllerInterface) {
	rg.GET("/user/:userId", userController.GetUserById)
	rg.POST("/user", userController.CreateUser)
}
