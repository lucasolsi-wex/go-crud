package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/internal/controller"
	"github.com/lucasolsi-wex/go-crud/internal/controller/routes"
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
		return
	}

	repo := repository.NewUserRepository(dbConnection)
	userService := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(viper.GetString("GIN_PORT")); err != nil {
		log.Fatal(err)
	}
}
