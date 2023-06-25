package handlers

import (
	"github.com/gin-gonic/gin"
	"handler/pkg/models"
	"net/http"
)

// ParseAndHandle Parses all data from port
func (handler *Handler) parse(ctx *gin.Context) {
	data, err := handler.services.Parser.ParseAllExchangers()
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
	return
}

func (handler *Handler) parseByCity(ctx *gin.Context) {
	var city = ctx.Param("city")
	data, err := handler.services.Parser.ParseExchangersByCity(city)

	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
	return
}

// database
func (handler *Handler) getEInfo(ctx *gin.Context) {
	data, err := handler.services.GetExchangersInfoTable()
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(200, data)
}

func (handler *Handler) getEKeys(ctx *gin.Context) {
	data, err := handler.services.GetExchangersKeysTable()
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(200, data)
}

func (handler *Handler) getECurrencies(ctx *gin.Context) {
	data, err := handler.services.GetExchangersCurrenciesTable()
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(200, data)
}
