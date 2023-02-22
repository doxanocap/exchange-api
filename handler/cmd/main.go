package main

import (
	"fmt"
	"handler/pkg/app"
	"handler/pkg/handlers"
	"handler/pkg/repository"
	"handler/pkg/repository/postgres"
	"handler/pkg/services"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	//url := fmt.Sprintf("http://localhost:8050/health")
	//
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//if err == nil {
	//	resp.Body.Close()
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//log.Println(string(body))
	//
	//return
	//
	server := app.Server{}

	// dev

	// db := postgres.InitDB(postgres.Config{
	// 	Host:     app.Public.ENV["POSTGRES_HOST"],
	// 	Port:     app.Public.ENV["POSTGRES_PORT"],
	// 	Username: app.Public.ENV["POSTGRES_USER"],
	// 	DBName:   app.Public.ENV["POSTGRES_DB"],
	// 	SSLMode:  app.Public.ENV["POSTGRES_SSL"],
	// 	Password: app.Public.ENV["POSTGRES_PASSWORD"]})

	// prod

	//domain := os.Getenv("DOMAINS.KZT_PARSER")

	app.InitEnvironment()

	db := postgres.InitDB(postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSL"),
		Password: os.Getenv("POSTGRES_PASSWORD")})

	if err := postgres.Migrations(db); err != nil {
		fmt.Printf("Migration went successfull:\nMigration file -> handler%s\n", err.Error())
		return
	}

	repos := repository.InitRepository(db)
	service := services.InitServices(repos)
	handler := handlers.InitHandler(service)

	if err := server.Run(app.Public.ENV["PORT"], handler.InitRoutes()); err != nil {
		log.Fatalf("Error whilte running http server: %s", err.Error())
	}
}
