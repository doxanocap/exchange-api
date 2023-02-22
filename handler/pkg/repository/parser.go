package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"handler/pkg/models"
	"handler/pkg/repository/postgres"
	"time"
)

type ParserModels struct {
	psql *sqlx.DB
	_m1  map[int]models.ExchangerKeys
	_m2  map[string]models.ExchangerKeys
}

func NewParserModels(db *sqlx.DB) *ParserModels {
	return &ParserModels{
		psql: db,
		_m1:  map[int]models.ExchangerKeys{},
		_m2:  map[string]models.ExchangerKeys{},
	}
}

func (parser *ParserModels) UpdateEInfoTableConst(eInfoData []models.ExchangerInfo) error {
	// Initializing to map to establish faster search from array
	tableDataMap := map[int]models.ExchangerInfo{}

	// Selected all existing exchanger_info from SQL table,
	// to make sure that we will not INSERT duplicate of any exchanger
	tableData, err := parser.selectEInfoData()
	if err != nil {
		return err
	}

	// tableDataMap is map of tableData -> tableDataMap[id]
	// Writing data from SQL table into map to make search faster
	for _, exchanger := range tableData {
		tableDataMap[exchanger.ExchangerId] = exchanger
	}

	// Finding out all exchangers that was parsed but not found in SQL table -> newEKeys
	var newEInfoData []models.ExchangerInfo
	for _, exchanger := range eInfoData {
		if _, ok := tableDataMap[exchanger.ExchangerId]; !ok {
			newEInfoData = append(newEInfoData, exchanger)
		}
	}

	// Inserting those new exchanger keys into sql table
	err = parser.insertEInfoTable(newEInfoData)
	return err
}

func (parser *ParserModels) UpdateEKeysTableConst(eKeys []models.ExchangerKeys) error {
	// Initializing to map to establish faster search from array
	tableDataMap := map[string]models.ExchangerKeys{}

	// Selected all existing EKeys from SQL table,
	// to make sure that we will not INSERT duplicate of any exchanger
	tableData, err := parser.selectEKeysData()
	if err != nil {
		return err
	}

	// tableDataMap is map of tableData -> tableDataMap[string]
	// Writing data from SQL table into map to make search faster
	for _, exchanger := range tableData {
		tableDataMap[exchanger.Name] = exchanger
	}

	// Finding out all exchangers that was parsed but not found in SQL table -> newEKeys
	var newEKeys []models.ExchangerKeys
	for _, exchanger := range eKeys {
		if _, ok := tableDataMap[exchanger.Name]; !ok {
			newEKeys = append(newEKeys, exchanger)
		}
	}

	// Inserting those new exchanger keys into sql table
	newTableEKeys, err := parser.insertIntoEKeysTable(newEKeys)
	if err != nil {
		return err
	}

	// Updating our TableKeys array with new values
	tableData = append(tableData, newTableEKeys...)
	for _, exchanger := range tableData {
		parser._m1[exchanger.Id] = exchanger
		parser._m2[exchanger.Name] = exchanger
	}
	return nil
}

func (parser *ParserModels) InsertKZTCurrencies(exchangers []models.ExchangerCurrencies) error {

	query := fmt.Sprintf(`
		INSERT INTO %s 
		(exchanger_id, usd_buy, usd_sell, eur_buy, eur_sell, rub_buy, rub_sell, updated_time)`,
		postgres.ExchangersCurrenciesTable)

	for _, exchanger := range exchangers {
		query += fmt.Sprintf(`
			(%d,%f,%f,%f,%f,%f,%f,%d)	
			`, exchanger.ExchangerId,
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

// GetKeysById && GetKeysByName: Methods for interaction with parser,_m1, parser._m2 maps in the higher levels
func (parser *ParserModels) GetKeysById(id int) models.ExchangerKeys {
	return parser._m1[id]
}

func (parser *ParserModels) GetKeysByName(name string) models.ExchangerKeys {
	return parser._m2[name]
}

//
//
// @Bellow are only private methods for this service
//
//

func (parser *ParserModels) insertEInfoTable(eInfoData []models.ExchangerInfo) error {
	// Inserting data into exchangers_info table
	if len(eInfoData) == 0 {
		return nil
	}

	query := fmt.Sprintf(`
		INSERT INTO %s 
		(exchanger_id, address, link, special_offer, phone_numbers)
		VALUES`, postgres.ExchangersInfoTable)

	// For loop to multiple insertion on single SLQ request
	for i, eInfo := range eInfoData {
		query += fmt.Sprintf(`
			(%d, '%s', '%s', '%s', '%s')`, eInfo.ExchangerId, eInfo.Address, eInfo.Link, eInfo.SpecialOffer, eInfo.PhoneNumbers)

		if i == len(eInfoData)-1 {
			query += ";"
			break
		}
		query += ", "
	}

	res, err := parser.psql.DB.Query(query)
	if err != nil {
		return err
	}
	err = res.Close()
	return err
}

func (parser *ParserModels) insertIntoEKeysTable(eKeys []models.ExchangerKeys) ([]models.ExchangerKeys, error) {
	if len(eKeys) == 0 {
		return nil, nil
	}
	var tableEKeys []models.ExchangerKeys

	query := fmt.Sprintf(`
		INSERT INTO %s (city, name) VALUES`, postgres.ExchangersKeysTable)

	for i, exchanger := range eKeys {
		query += fmt.Sprintf("('%s','%s')", exchanger.City, exchanger.Name)
		if i == len(eKeys)-1 {
			query += " RETURNING *;"
			break
		}
		query += ", \n"
	}

	if err := parser.psql.Select(&tableEKeys, query); err != nil {
		return nil, err
	}
	return tableEKeys, nil
}

func (parser *ParserModels) selectEInfoData() ([]models.ExchangerInfo, error) {
	var eInfoData []models.ExchangerInfo
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ExchangersInfoTable)
	err := parser.psql.Select(&eInfoData, query)
	if err != nil {
		return nil, err
	}
	return eInfoData, nil
}

func (parser *ParserModels) selectEKeysData() ([]models.ExchangerKeys, error) {
	var eKeys []models.ExchangerKeys
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ExchangersKeysTable)
	if err := parser.psql.Select(&eKeys, query); err != nil {
		return nil, err
	}
	return eKeys, nil
}
