package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	server "github.com/trucktrace"
	"github.com/trucktrace/internal/repository"
	service "github.com/trucktrace/internal/services"
	"github.com/trucktrace/pkg/database"
)

var postgresClient *gorm.DB
var ser *server.Server

func main() {

	mongoClient, err := database.InitMongoDB(&database.MongoConfig{
		Username: "trucktrace",
		Password: "trucktrace",
		Host:     "localhost",
		Port:     "27017",
	})

	if err != nil {
		log.Fatal("Can not init mongo")
	}

	postgresClient, pgerr := database.InitPostgresDB()

	if pgerr != nil {
		log.Fatal("can not init postgres")
	}
	defer postgresClient.Close()

	repos := repository.NewRepository(mongoClient, postgresClient)
	service := service.NewService(repos)
	handlers := controllers.NewHandler(service)

	if err := http.ListenAndServe("localhost:8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Can't run server... %s", err.Error())
	}

	initiateRoutes := ser.Run("8080", controllers.InitRoutes())

	if initiateRoutes != nil {
		// logger.FatalLogger("Can't initiate routes", "cmd/main/main.go", initiateRoutes.Error())
		ser.ShutDown(context.Background())
	}

	// Printing the starting message
	// logger.InfoLogger("Server was running on 127.0.0.1:8080", "cmd/main/main.go", "Success")
}
