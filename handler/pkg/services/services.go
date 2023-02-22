package services

import (
	"handler/pkg/models"
	"handler/pkg/repository"
)

type Dispatcher interface {
	Data() models.ParserResponse
}

type Parser interface {
	InsertKZTCurrencies() error
	GetAllExchangers() ([]models.ParserResponse, error)
	GetExchangersByCity(city string) ([]models.ParserResponse, error)
}

type Request interface {
}

type Services struct {
	Dispatcher
	Parser
	Request
}

func InitServices(repository *repository.Repository) *Services {
	return &Services{
		Parser: InitParserService(repository.Parser),
	}
}
