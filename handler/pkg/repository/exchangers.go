package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"handler/pkg/models"
	"handler/pkg/repository/postgres"
)

type ExchangersModels struct {
	psql *sqlx.DB
}

func NewExchangersModels(db *sqlx.DB) *ExchangersModels {
	return &ExchangersModels{psql: db}
}

type CurrenciesData struct {
	UploadTime uint64  `json:"upload_time" db:"upload_time"`
	USD_BUY    float32 `json:"USD_BUY" db:"usd_buy"`
	USD_SELL   float32 `json:"USD_SELL" db:"usd_sell"`
	EUR_BUY    float32 `json:"EUR_BUY" db:"eur_buy"`
	EUR_SELL   float32 `json:"EUR_SELL" db:"eur_sell"`
	RUB_BUY    float32 `json:"RUB_BUY" db:"rub_buy"`
	RUB_SELL   float32 `json:"RUB_SELL" db:"rub_sell"`
}

func (exchangers *ExchangersModels) SelectExchangersData(params models.ExchangerInfoParams) ([]models.ExchangerData, error) {
	var data []models.ExchangerData
	query := fmt.Sprintf(`
			SELECT 
				k.id, 
				k.city, 
				k.name, 
				i.address, 
				i.wholesale, 
				i.updated_time, 
				i.phone_numbers 
			FROM %s k
				INNER JOIN %s i
			ON k.id = i.exchanger_id`,
		postgres.ExchangersKeysTable, postgres.ExchangersInfoTable)
	if params.Name != "" {
		query += fmt.Sprintf(" WHERE k.name = '%s'", params.Name)
	} else if params.City != "" {
		query += fmt.Sprintf(" WHERE k.city = '%s'", params.City)
	}

	if params.Wholesale {
		query += fmt.Sprintf(" AND LENGTH(i.wholesale) != 0")
	}

	query += fmt.Sprintf(";")
	log.Println(query)
	if err := exchangers.psql.Select(&data, query); err != nil {
		return nil, err
	}
	return data, nil

}

func (exchangers *ExchangersModels) SelectCurrenciesData(params models.CurrenciesDataParams) ([]CurrenciesData, error) {
	var data []CurrenciesData
	query := fmt.Sprintf(`
		SELECT 
			upload_time,
			avg(usd_buy) as usd_buy,
			avg(usd_sell) as usd_sell,
			avg(eur_buy) as eur_buy,
			avg(eur_sell) as eur_sell,
			avg(rub_buy) as rub_buy,
			avg(rub_sell) as rub_sell
		FROM %s
			WHERE usd_buy > 0 AND upload_time >= %d AND upload_time <= %d
		GROUP BY upload_time
			ORDER BY upload_time ASC;
		`, postgres.ExchangersCurrenciesTable, params.From, params.To)
	if err := exchangers.psql.Select(&data, query); err != nil {
		return nil, err
	}
	return data, nil
}
