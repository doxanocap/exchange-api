package router

import (
	"fmt"
	"net/http"
	"webchat/pkg/repository/models"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) webSocketConn(ctx *gin.Context) {
	header := http.Header{}
	conn, err := ConnUpgrade.Upgrade(ctx.Writer, ctx.Request, header)
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{
			Status:  500,
			Message: fmt.Sprintf("upgrading connection: %s", err.Error()),
		})
	}

	pool := handler.services.Pool
	client := handler.services.Client.NewClient(ctx, pool, conn)

	go client.Reader()
	go client.Writer()
}
