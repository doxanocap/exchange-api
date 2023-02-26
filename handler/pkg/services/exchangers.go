package services

import (
	"handler/pkg/models"
	"handler/pkg/repository"
)

type ExchangersService struct {
	exchangerModels repository.Exchanger
	parserModels    repository.Parser
}

func InitExchangersServices(repo *repository.Repository) *ExchangersService {
	return &ExchangersService{
		exchangerModels: repo.Exchanger,
		parserModels:    repo.Parser,
	}
}

func (exchanger *ExchangersService) GetExchangersData(params models.ExchangerInfoParams) ([]models.ExchangerData, error) {
	response, err := exchanger.exchangerModels.SelectExchangersData(params)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (exchanger *ExchangersService) GetCurrenciesData(params models.CurrenciesDataParams) ([]models.CurrenciesData, error) {
	data, err := exchanger.exchangerModels.SelectCurrenciesData(params)
	if err != nil {
		return nil, err
	}

	var response []models.CurrenciesData
	for _, e := range data {
		response = append(response, models.CurrenciesData{
			UploadTime: e.UploadTime,
			Currencies: models.Currencies{
				USD: [2]float32{e.USD_BUY, e.USD_SELL},
				EUR: [2]float32{e.EUR_BUY, e.EUR_SELL},
				RUB: [2]float32{e.RUB_BUY, e.RUB_SELL},
			},
		})
	}
	return response, nil
}
