package handlers

import (
	"handler/pkg/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func InitHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/healthcheck", handler.healthcheck)
	{
		parser := router.Group("/kzt-parser")
		{
			parser.GET("/parse", handler.parse)
			parser.GET("/parse/:city", handler.parseByCity)

			database := parser.Group("/database")
			{
				database.GET("/exchanger-info", handler.getEInfo)
				database.GET("/exchanger-keys", handler.getEKeys)
				database.GET("/exchanger-currencies", handler.getECurrencies)
			}
		}

		data := router.Group("/data")
		{
			data.POST("/exchangers", handler.getExchangersData)
			data.POST("/currencies", handler.getCurrenciesData)
		}

	}
	return router
}
