package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"handler/pkg/app"
	"handler/pkg/repository/postgres"
	"strings"
	"time"
)

type ParserModels struct {
	psql *sqlx.DB
	_m1  map[int]app.ExchangerKeys
	_m2  map[string]app.ExchangerKeys
}

func NewParserModels(db *sqlx.DB) *ParserModels {
	return &ParserModels{
		psql: db,
		_m1:  map[int]app.ExchangerKeys{},
		_m2:  map[string]app.ExchangerKeys{},
	}
}

func (parser *ParserModels) UpdateExchangersTableConst(exchangers []app.Exchanger) error {
	err := parser.updateEKeysTableConst(exchangers)
	if err != nil {
		return err
	}
	err = parser.updateEInfoTableConst(exchangers)
	return err
}

func (parser *ParserModels) InsertKZTCurrencies(exchangers []app.Exchanger) error {
	query := fmt.Sprintf(`
		INSERT INTO %s 
		(exchanger_id, usd_buy, usd_sell, eur_buy, eur_sell, rub_buy, rub_sell, updated_time)`,
		postgres.ExchangersCurrenciesTable)

	for _, exchanger := range exchangers {
		query += fmt.Sprintf(`
			(%d,%f,%f,%f,%f,%f,%f, %d)	
			`, exchanger.Id,
			exchanger.USD_BUY, exchanger.USD_SELL,
			exchanger.EUR_BUY, exchanger.EUR_SELL,
			exchanger.RUB_BUY, exchanger.RUB_SELL,
			time.Now().Unix())
	}

	res, err := parser.psql.DB.Query(query)
	if err != nil {
		return err
	}
	return res.Close()
}

//
//
// @Bellow are only private methods for this service
//
//

func (parser *ParserModels) updateEInfoTableConst(exchangers []app.Exchanger) error {
	var eInfo []app.ExchangerInfo
	for _, exchanger := range exchangers {
		info := app.ExchangerInfo{
			Id:           parser._m2[exchanger.Name].Id,
			Address:      exchanger.Address,
			Link:         exchanger.Link,
			SpecialOffer: exchanger.SpecialOffer,
			PhoneNumbers: exchanger.PhoneNumbers,
		}
		eInfo = append(eInfo, info)
	}

	err := parser.insertEInfo(eInfo)
	return err
}

func (parser *ParserModels) updateEKeysTableConst(eKeys []app.Exchanger) error {
	// Initializing to map to establish faster search from array
	var tableDataMap map[string]app.ExchangerKeys

	// Selected all existing EKeys from SQL table, to make sure that we will not INSERT duplicate of any exchanger
	tableEKeys, err := parser.selectEKeys()
	if err != nil {
		return err
	}

	// Writing data from SQL table into map to make search faster
	for _, exchanger := range tableEKeys {
		tableDataMap[exchanger.Name] = exchanger
	}

	// Finding out all exchangers that was parsed but not found in SQL table -> newEKeys
	var newEKeys []app.ExchangerKeys
	for _, exchanger := range eKeys {
		if _, ok := tableDataMap[exchanger.Name]; !ok {
			newEKeys = append(newEKeys, app.ExchangerKeys{
				Id:   exchanger.Id,
				City: exchanger.City,
				Name: exchanger.Name})
		}
	}

	// Inserting those new exchanger keys into sql table
	newTableEKeys, err := parser.insertEKeys(newEKeys)
	if err != nil {
		return err
	}

	// Updating our TableKeys array with new values
	tableEKeys = append(tableEKeys, newTableEKeys...)
	for _, exchanger := range tableEKeys {
		parser._m1[exchanger.Id] = exchanger
		parser._m2[exchanger.Name] = exchanger
	}
	return nil
}

func (parser *ParserModels) insertEInfo(eInformation []app.ExchangerInfo) error {
	// Inserting data with ON CONFLICT condition
	// to update table only if data with particular ID
	// does not exist in the table.

	query := fmt.Sprintf(`
		INSERT INTO %s 
		(exchanger_id, address, link, special_offer, phones)
	`, postgres.ExchangersInfoTable)

	// For loop to multiple insertion on single SLQ request
	for _, eInfo := range eInformation {
		query += fmt.Sprintf(`
			(%d, %s, %s, %s, %s)
		`, eInfo.Id, eInfo.Address, eInfo.Link, eInfo.SpecialOffer, strings.Join(eInfo.PhoneNumbers, ","))
	}

	// ON CONFLICT update table
	query += fmt.Sprintf(`
		ON CONFLICT 
		(exchanger_id) 
		DO UPDATE 
		SET address = excluded.name,
			link = excluded.link,
			special_offer = excluded.special_offer,
			phones = excluded.phones`)

	res, err := parser.psql.DB.Query(query)
	if err != nil {
		return err
	}
	err = res.Close()
	return err
}

func (parser *ParserModels) insertEKeys(eKeys []app.ExchangerKeys) ([]app.ExchangerKeys, error) {
	var tableEKeys []app.ExchangerKeys
	query := fmt.Sprintf("INSERT INTO %s (city, name) ", postgres.ExchangersKeysTable)
	for i, exchanger := range eKeys {
		query += fmt.Sprintf("(%s,%s)", exchanger.City, exchanger.Name)
		if i == len(eKeys)-1 {
			query += "RETURNING *;"
			break
		}
		query += ","
	}
	if err := parser.psql.Select(&tableEKeys, query); err != nil {
		return nil, err
	}
	return tableEKeys, nil
}

func (parser *ParserModels) selectEKeys() ([]app.ExchangerKeys, error) {
	var eKeys []app.ExchangerKeys
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ExchangersKeysTable)
	if err := parser.psql.Select(&eKeys, query); err != nil {
		return nil, err
	}
	return eKeys, nil
}
