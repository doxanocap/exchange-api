package services

import (
	"handler/pkg/app"
	"handler/pkg/repository"
)

type Dispatcher interface {
	Data() app.ParserResponse
}

type Parser interface {
	InsertKZTCurrencies() error
	GetAllExchangers() ([]app.ParserResponse, error)
	GetExchangersByCity(city string) ([]app.ParserResponse, error)
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
		//Parser: InitParserService(repository.Parser),
	}
}
