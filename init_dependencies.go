package main

import (
	"github.com/lucasolsi-wex/go-crud/src/controller"
	"github.com/lucasolsi-wex/go-crud/src/model/service"
	"github.com/lucasolsi-wex/go-crud/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	domainService := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(domainService)
}
