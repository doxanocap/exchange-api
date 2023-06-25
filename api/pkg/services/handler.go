package services

import (
	"api/pkg/configs"
	"api/pkg/dispatcher"
	"api/pkg/models"
	"encoding/json"
)

var (
	HANDLER_URL = configs.ENV("DOMAINS_HANDLER") + "/data"
)

type HandlerService struct {
	dp *dispatcher.NewDispatcher
}

func NewHandlerService(dp *dispatcher.NewDispatcher) *HandlerService {
	return &HandlerService{dp}
}

func (service *HandlerService) ExchangersData(body []byte) ([]models.ExchangersResponse, models.ErrorResponse) {
	res, err := service.dp.Post("/data/exchangers", body)
	if err.Message != "" {
		return []models.ExchangersResponse{}, models.ErrorResponse{
			Status:  500,
			Message: err.Message}
	}

	var _err models.ErrorResponse
	if err := json.Unmarshal(res, &_err); err == nil {
		return []models.ExchangersResponse{}, models.ErrorResponse{
			Status:  _err.Status,
			Message: "ERROR response from handler/data/exchangers -> " + _err.Message,
		}
	}

	var exchangers []models.ExchangersResponse
	if err := json.Unmarshal(res, &exchangers); err != nil {
		return []models.ExchangersResponse{}, models.ErrorResponse{
			Status:  200,
			Message: "ERROR while unmarshal handler/data/exchangers -> " + err.Error(),
		}
	}
	return exchangers, models.ErrorResponse{}
}

func (service *HandlerService) CurrenciesData(body []byte) ([]models.CurrenciesResponse, models.ErrorResponse) {
	res, err := service.dp.Post("/data/currencies", body)
	if err.Message != "" {
		return []models.CurrenciesResponse{}, models.ErrorResponse{
			Status:  500,
			Message: err.Message}
	}

	var _err models.ErrorResponse
	if err := json.Unmarshal(res, &_err); err == nil {
		return []models.CurrenciesResponse{}, models.ErrorResponse{
			Status:  _err.Status,
			Message: "ERROR response from handler/data/currencies -> " + _err.Message,
		}
	}

	var currencies []models.CurrenciesResponse
	if err := json.Unmarshal(res, &currencies); err != nil {
		return []models.CurrenciesResponse{}, models.ErrorResponse{
			Status:  200,
			Message: "ERROR while unmarshal handler/data/currencies -> " + err.Error(),
		}
	}

	return currencies, models.ErrorResponse{}
}
