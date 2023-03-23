package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"webchat/pkg/services"
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

	webchat := router.Group("/web-chat")
	{
		pool := webchat.Group("/pool")
		{
			pool.GET("/", handler.webSocketConn)
		}

		user := webchat.Group("/user")
		{
			user.GET("/online")
		}
		{
			webchat.POST("/active-chats")
			webchat.GET("/messages/:chatUID")
		}
	}
	return router
}
