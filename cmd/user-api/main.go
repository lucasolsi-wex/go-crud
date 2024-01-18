package main

import (
	"context"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/internal/controller"
	"github.com/lucasolsi-wex/go-crud/internal/database"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/lucasolsi-wex/go-crud/internal/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	gin.SetMode(viper.GetString("GIN_MODE"))

	if err != nil {
		return
	}

	fmt.Println()

	dbConnection, err := database.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal("Error while establishing connection to database: ", err)
	}

	repo := repository.NewUserRepository(dbConnection)
	userService := service.UserInterfaceService{Repository: repo}
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/user/:userId", userController.FindUserById)
		v1.POST("/user", userController.CreateUser)
	}

	if err := endless.ListenAndServe(viper.GetString("GIN_PORT"), router); err != nil {
		log.Fatal(err)
	}
}
