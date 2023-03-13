package router

import (
	"api/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (router *Router) ExchangersData(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		router.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	res, _err := router.services.Handler.ExchangersData(data)
	if _err.Message != "" {
		router.newErrorResponse(ctx, _err)
		return
	}
	ctx.JSON(200, res)
}

func (router *Router) CurrenciesData(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		router.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	res, err1 := router.services.Handler.CurrenciesData(data)
	log.Println(res, err1)

	if err1.IsError() {
		router.newErrorResponse(ctx, err1)
		return
	}
	ctx.JSON(200, res)
}
