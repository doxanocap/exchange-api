package cmd

import (
	"context"
	router "core/internal/rest"
	"core/pkg/config"
	"core/pkg/logger"
	"core/pkg/postgres"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"
)

func Run() {
	config.InitConf()
	logger.Init(viper.GetBool("isProd"), viper.GetBool("isJson"))

	conn, err := postgres.Connect(context.Background(), viper.GetString("pg_dsn"))
	if err != nil {
		logger.Log.Fatal("connection to postgres ->", zap.Error(err))
	}

	app := InitApp(conn)

	db := postgres.InitDB(postgres.Config{
		Host:     configs.ENV("POSTGRES_HOST"),
		Port:     configs.ENV("POSTGRES_PORT"),
		Username: configs.ENV("POSTGRES_USER"),
		DBName:   configs.ENV("POSTGRES_DB"),
		SSLMode:  configs.ENV("POSTGRES_SSL"),
		Password: configs.ENV("POSTGRES_PASSWORD")})

	if err := postgres.Migrations(db); err != nil {
		log.Printf("Migration went successfull: Migration file -> rest%s\n", err.Error())
		return
	}

	dp := dispatcher.InitDispatcher()
	service := services.InitServices(dp)
	handler := router.InitRouter(service)

	if err := server.Run(configs.ENV("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error whilte running httpServer httpServer: %s", err.Error())
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
