package handlers

import (
	"api/pkg/services"
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
		api := router.Group("/api")
		{
			auth := api.Group("/auth")
			{
				auth.POST("/sign-in", handler.SignIn)
				auth.POST("/sign-up", handler.SignIn)
				auth.GET("/sign-out")
				auth.GET("/refresh-token")
			}
		}
	}
	return router
}
