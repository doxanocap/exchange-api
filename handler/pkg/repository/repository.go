package repository

import (
	"github.com/jmoiron/sqlx"
	"handler/pkg/app"
)

type Dispatcher interface {
}

type Parser interface {
	UpdateExchangersTableConst(exchangers []app.Exchanger) error
	InsertKZTCurrencies(exchangers []app.Exchanger) error
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
