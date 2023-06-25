package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"webchat/pkg/repository"
)

type Pool interface {
	HandleWebSocketConn()
}

type Chat interface {
}

type Client interface {
	NewClient(context *gin.Context, pool *PoolService, conn *websocket.Conn) *ClientType
}

type Services struct {
	Client *ClientService
	Pool   *PoolService
	Chat   *ChatService
}

func InitServices(repository *repository.Repository) *Services {
	return &Services{
		Client: NewClientService(repository),
		Pool:   NewPoolService(repository),
		Chat:   NewChatServices(repository),
	}
}
