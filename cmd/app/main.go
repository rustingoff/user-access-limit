package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/rustingoff/user-access-limit/internal/handler"
	"github.com/rustingoff/user-access-limit/internal/repository"
	"github.com/rustingoff/user-access-limit/internal/service"
	"github.com/rustingoff/user-access-limit/pkg/database"
)

var postgresClient *gorm.DB

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

	repos := repository.NewRepository(mongoClient, postgresClient)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	if err := http.ListenAndServe("localhost:8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Can't run server... %s", err.Error())
	}

	log.Print("Server was started")
}
