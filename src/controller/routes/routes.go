package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/controller"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/user/:userId", controller.GetUserById)
	rg.POST("/user", controller.CreateUser)
}
