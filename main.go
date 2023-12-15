package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/database"
	"github.com/lucasolsi-wex/go-crud/src/controller"
	"github.com/lucasolsi-wex/go-crud/src/controller/routes"
	"github.com/lucasolsi-wex/go-crud/src/model/service"
	"github.com/lucasolsi-wex/go-crud/src/repository"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		return
	}

	fmt.Println()

	dbConnection, err := database.NewMongoDBConnection(context.Background())
	if err != nil {
		return
	}

	repo := repository.NewUserRepository(dbConnection)
	userService := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
