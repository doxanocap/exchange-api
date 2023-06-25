package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"webchat"
	"webchat/pkg/configs"
	"webchat/pkg/handlers"
	"webchat/pkg/repository"
	"webchat/pkg/repository/postgres"
	"webchat/pkg/services"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	configs.InitEnvironment()
	server := webchat.Server{}

	db := postgres.InitDB(postgres.Config{
		Host:     configs.ENV("POSTGRES_HOST"),
		Port:     configs.ENV("POSTGRES_PORT"),
		Username: configs.ENV("POSTGRES_USER"),
		DBName:   configs.ENV("POSTGRES_DB"),
		SSLMode:  configs.ENV("POSTGRES_SSL"),
		Password: configs.ENV("POSTGRES_PASSWORD")})

	if err := postgres.Migrations(db); err != nil {
		log.Printf("Migration went successfull: Migration file -> handler%s\n", err.Error())
		return
	}

	repo := repository.InitRepository(db)
	service := services.InitServices(repo)
	handler := handlers.InitHandler(service)

	if err := server.Run(configs.ENV("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error whilte running http server: %s", err.Error())
	}

	log.Printf("Messanger service started \n")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Messanger service shut down\n")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("service shutting down error - %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("Messanger service: database shutting down error - %s", err.Error())
	}

}
