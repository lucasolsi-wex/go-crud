package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasolsi-wex/go-crud/src/config/database"
	"github.com/lucasolsi-wex/go-crud/src/controller/routes"
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

	userController := initDependencies(dbConnection)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
