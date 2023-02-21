package repository

import (
	"github.com/jmoiron/sqlx"
	"handler/pkg/app"
)

type Dispatcher interface {
}

type Parser interface {
	UpdateEInfoTableConst(eInfo []app.ExchangerInfo) error
	UpdateEKeysTableConst(eKeys []app.ExchangerKeys) error
	InsertKZTCurrencies(exchangers []app.ExchangerCurrencies) error
	GetKeysById(id int) app.ExchangerKeys
	GetKeysByName(name string) app.ExchangerKeys
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
