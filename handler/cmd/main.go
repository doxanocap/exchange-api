package main

import (
	"context"
	"fmt"
	app "handler"
	"handler/pkg/configs"
	"handler/pkg/handlers"
	"handler/pkg/repository"
	"handler/pkg/repository/postgres"
	"handler/pkg/services"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
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

	log.Printf("Handler service started \n")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Handler service shut down\n")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("service shutting down error - %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("handler service: database shutting down error - %s", err.Error())
	}
}
