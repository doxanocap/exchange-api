package services

import (
	"handler/pkg/models"
	"handler/pkg/repository"
)

type Exchangers interface {
	GetExchangersData(params models.ExchangerInfoParams) ([]models.ExchangerData, error)
	GetCurrenciesData(params models.CurrenciesDataParams) ([]models.CurrenciesData, error)
}

type Parser interface {
	ParseAllExchangers() ([]models.ParserResponse, error)
	ParseExchangersByCity(city string) ([]models.ParserResponse, error)

	GetExchangersKeysTable() ([]repository.ExchangerKeys, error)
	GetExchangersInfoTable() ([]repository.ExchangerInfo, error)
	GetExchangersCurrenciesTable() ([]models.ExchangerCurrenciesResponse, error)
}

type Services struct {
	Exchangers
	Parser
}

func InitServices(repository *repository.Repository) *Services {
	return &Services{
		Exchangers: InitExchangersServices(repository),
		Parser:     InitParserService(repository.Parser),
	}
}
