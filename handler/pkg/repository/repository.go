package repository

import (
	"github.com/jmoiron/sqlx"
	"handler/pkg/models"
)

type Exchanger interface {
	SelectExchangersData(params models.ExchangerInfoParams) ([]models.ExchangerData, error)
	SelectCurrenciesData(params models.CurrenciesDataParams) ([]CurrenciesData, error)
}

type Parser interface {
	GetKeysById(id int) ExchangerKeys
	GetKeysByName(name string) ExchangerKeys

	SelectExchangersInfo() ([]ExchangerInfo, error)
	SelectExchangersKeys() ([]ExchangerKeys, error)
	SelectExchangersCurrencies() ([]ExchangerCurrencies, error)

	UpdateEInfoTableConst(eInfo []ExchangerInfo) error
	UpdateEKeysTableConst(eKeys []ExchangerKeys) error

	InsertKZTCurrencies(exchangers []ExchangerCurrencies) error
	InsertEInfoTable(eInfoData []ExchangerInfo) error
	InsertIntoEKeysTable(eKeys []ExchangerKeys) ([]ExchangerKeys, error)
}

type Repository struct {
	Exchanger
	Parser
}

func InitRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Exchanger: NewExchangersModels(db),
		Parser:    NewParserModels(db),
	}
}
