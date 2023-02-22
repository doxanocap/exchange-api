package repository

import (
	"github.com/jmoiron/sqlx"
	"handler/pkg/models"
)

type Dispatcher interface {
}

type Parser interface {
	UpdateEInfoTableConst(eInfo []models.ExchangerInfo) error
	UpdateEKeysTableConst(eKeys []models.ExchangerKeys) error
	InsertKZTCurrencies(exchangers []models.ExchangerCurrencies) error
	GetKeysById(id int) models.ExchangerKeys
	GetKeysByName(name string) models.ExchangerKeys
}

type Request interface {
}

type Repository struct {
	Dispatcher
	Parser *ParserModels
	Request
}

func InitRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Parser: NewParserModels(db),
	}
}
