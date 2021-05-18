package main

import (
	"log"
	"net/http"

	"github.com/rustingoff/user-access-limit/internal/handler"
	"github.com/rustingoff/user-access-limit/internal/repository"
	"github.com/rustingoff/user-access-limit/internal/service"

	"github.com/jmoiron/sqlx"
)

func main() {

	var db *sqlx.DB

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	if err := http.ListenAndServe(":8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Can't run server... %s", err.Error())
	}
}
