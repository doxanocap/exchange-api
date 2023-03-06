package handlers

import (
	"api/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) newErrorResponse(ctx *gin.Context, err models.ErrorResponse) {
	logrus.Error(err.Message)
	ctx.AbortWithStatusJSON(err.Status, gin.H{"status": err.Status, "message": err.Message})
}

func (handler *Handler) healthcheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "handler service is alive"})
}
