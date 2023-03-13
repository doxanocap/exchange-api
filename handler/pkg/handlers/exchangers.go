package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"handler/pkg/models"
	"net/http"
)

type ExchangersByCityRequest struct {
	HasSale  bool  `json:"has_sale"`
	DataFrom int64 `json:"data_from"`
	DataTo   int64 `json:"data_to"`
}

// exchangers
func (handler *Handler) getExchangersData(ctx *gin.Context) {
	var params models.ExchangerInfoParams

	if err := ctx.BindJSON(&params); err != nil {
		log.Println("weqwewq ", err)
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := handler.services.Exchangers.GetExchangersData(params)
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (handler *Handler) getCurrenciesData(ctx *gin.Context) {
	var params models.CurrenciesDataParams
	if err := ctx.BindJSON(&params); err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := handler.services.Exchangers.GetCurrenciesData(params)
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
