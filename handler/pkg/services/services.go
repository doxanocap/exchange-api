package services

import (
	"handler/pkg/app"
	"handler/pkg/repository"
)

type Dispatcher interface {
	Data() app.Exchanger
}

type Parser interface {
	InsertKZTCurrencies() error
	GetAllExchangers() ([]app.Exchanger, error)
	GetExchangersByCity(city string) ([]app.Exchanger, error)
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
