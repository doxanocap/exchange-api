package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	app "handler"
	"handler/pkg/configs"
	"handler/pkg/handlers"
	"handler/pkg/repository"
	"handler/pkg/repository/postgres"
	"handler/pkg/services"
)

func main() {
	fmt.Printf("\n\n")
	configs.InitEnvironment()
	server := app.Server{}

	db := postgres.InitDB(postgres.Config{
		Host:     configs.ENV("POSTGRES_HOST"),
		Port:     configs.ENV("POSTGRES_PORT"),
		Username: configs.ENV("POSTGRES_USER"),
		DBName:   configs.ENV("POSTGRES_DB"),
		SSLMode:  configs.ENV("POSTGRES_SSL"),
		Password: configs.ENV("POSTGRES_PASSWORD")})

	if err := postgres.Migrations(db); err != nil {
		fmt.Printf("Migration went successfull: Migration file -> handler%s\n", err.Error())
		return
	}

	repos := repository.InitRepository(db)
	service := services.InitServices(repos)
	handler := handlers.InitHandler(service)

	if err := server.Run(configs.ENV("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error whilte running http server: %s", err.Error())
	}
}
