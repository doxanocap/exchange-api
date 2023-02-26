package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"handler/pkg/models"
)

func (handler *Handler) newErrorResponse(ctx *gin.Context, err models.ErrorResponse) {
	logrus.Error(err.Message)
	ctx.AbortWithStatusJSON(err.Status, err.Message)
}
