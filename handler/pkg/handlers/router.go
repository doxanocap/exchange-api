package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	parser := router.Group("/kzt-parser")
	{
		parser.GET("/parse", handler.ParseAndHandle)
	}

	api := router.Group("/api")
	{
		data := api.Group("/data")
		{
			data.GET("/", handler.data)
			data.GET("/:city", handler.dataByCity)
			data.GET("/:city/:name", handler.dataByExchanger)
		}

		exchangers := parser.Group("/exchangers")
		{
			exchangers.GET("/", handler.exchangers)
			exchangers.GET("/:city", handler.exchangersByCity)
			exchangers.GET("/:city/:name", handler.exchangersByName)
		}

		avg := parser.Group("/avg/currencies")
		{
			avg.GET("/:city", handler.avgCurrenciesByCity)
			avg.GET("/:city/:name", handler.avgCurrenciesByExchanger)
		}
	}
	return router
}
