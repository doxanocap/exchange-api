package service

import (
	"api/pkg/configs"
	"api/pkg/dispatcher"
	"api/pkg/repository"
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

func (service *HandlerService) ExchangersData(body []byte) ([]repository.ExchangersResponse, repository.ErrorResponse) {
	res, err := service.dp.Post("/data/exchangers", body)
	if err.Message != "" {
		return []repository.ExchangersResponse{}, repository.ErrorResponse{
			Status:  500,
			Message: err.Message}
	}

	var _err repository.ErrorResponse
	if err := json.Unmarshal(res, &_err); err == nil {
		return []repository.ExchangersResponse{}, repository.ErrorResponse{
			Status:  _err.Status,
			Message: "ERROR response from rest/data/exchangers -> " + _err.Message,
		}
	}

	var exchangers []repository.ExchangersResponse
	if err := json.Unmarshal(res, &exchangers); err != nil {
		return []repository.ExchangersResponse{}, repository.ErrorResponse{
			Status:  200,
			Message: "ERROR while unmarshal rest/data/exchangers -> " + err.Error(),
		}
	}
	return exchangers, repository.ErrorResponse{}
}

func (service *HandlerService) CurrenciesData(body []byte) ([]repository.CurrenciesResponse, repository.ErrorResponse) {
	res, err := service.dp.Post("/data/currencies", body)
	if err.Message != "" {
		return []repository.CurrenciesResponse{}, repository.ErrorResponse{
			Status:  500,
			Message: err.Message}
	}

	var _err repository.ErrorResponse
	if err := json.Unmarshal(res, &_err); err == nil {
		return []repository.CurrenciesResponse{}, repository.ErrorResponse{
			Status:  _err.Status,
			Message: "ERROR response from rest/data/currencies -> " + _err.Message,
		}
	}

	var currencies []repository.CurrenciesResponse
	if err := json.Unmarshal(res, &currencies); err != nil {
		return []repository.CurrenciesResponse{}, repository.ErrorResponse{
			Status:  200,
			Message: "ERROR while unmarshal rest/data/currencies -> " + err.Error(),
		}
	}

	return currencies, repository.ErrorResponse{}
}
