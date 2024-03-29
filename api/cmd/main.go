package main

import (
	app "api"
	"api/pkg/configs"
	"api/pkg/dispatcher"
	"api/pkg/router"
	"api/pkg/services"
	"context"
	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	configs.InitEnvironment()
	server := app.Server{}

	dp := dispatcher.InitDispatcher()
	service := services.InitServices(dp)
	handler := router.InitRouter(service)

	if err := server.Run(configs.ENV("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error whilte running http server: %s", err.Error())
	}

	log.Printf("Handler service started \n")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := dp.ServicesShutDownCheck(); err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("API gateway has been shut down\n")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("service shutting down error - %s", err.Error())
	}
}
