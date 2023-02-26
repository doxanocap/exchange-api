package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"handler/pkg/repository/postgres"
	"time"
)

type ParserModels struct {
	psql *sqlx.DB
	_m1  map[int]ExchangerKeys
	_m2  map[string]ExchangerKeys
}

type ExchangerKeys struct {
	Id   int    `json:"id" db:"id"`
	City string `json:"city" db:"city"`
	Name string `json:"name" db:"name"`
}

type ExchangerInfo struct {
	ExchangerId  int    `json:"exchanger_id" db:"exchanger_id"`
	Address      string `json:"address" db:"address"`
	WholeSale    string `json:"wholesale" db:"wholesale"`
	UpdatedTime  uint64 `json:"updated_time" db:"updated_time"`
	PhoneNumbers string `json:"phone_numbers" db:"phone_numbers"`
}

type ExchangerCurrencies struct {
	ExchangerId int     `json:"-" db:"exchanger_id"`
	UploadTime  uint64  `json:"upload_time" db:"upload_time"`
	USD_BUY     float32 `json:"USD_BUY" db:"usd_buy"`
	USD_SELL    float32 `json:"USD_SELL" db:"usd_sell"`
	EUR_BUY     float32 `json:"EUR_BUY" db:"eur_buy"`
	EUR_SELL    float32 `json:"EUR_SELL" db:"eur_sell"`
	RUB_BUY     float32 `json:"RUB_BUY" db:"rub_buy"`
	RUB_SELL    float32 `json:"RUB_SELL" db:"rub_sell"`
}

func NewParserModels(db *sqlx.DB) *ParserModels {
	return &ParserModels{
		psql: db,
		_m1:  map[int]ExchangerKeys{},
		_m2:  map[string]ExchangerKeys{},
	}
}

func (parser *ParserModels) UpdateEKeysTableConst(eKeys []ExchangerKeys) error {
	// Initializing to map to establish faster search from array
	tableDataMap := map[string]ExchangerKeys{}

	// Selected all existing EKeys from SQL table,
	// to make sure that we will not INSERT duplicate of any exchanger
	tableData, err := parser.SelectExchangersKeys()
	if err != nil {
		return err
	}

	// tableDataMap is map of tableData -> tableDataMap[string]
	// Writing data from SQL table into map to make search faster
	for _, exchanger := range tableData {
		tableDataMap[exchanger.Name] = exchanger
	}

	// Finding out all exchangers that was parsed but not found in SQL table -> newEKeys
	var newEKeys []ExchangerKeys
	for _, exchanger := range eKeys {
		if _, ok := tableDataMap[exchanger.Name]; !ok {
			newEKeys = append(newEKeys, exchanger)
		}
	}

	// Inserting those new exchanger keys into sql table
	newTableEKeys, err := parser.InsertIntoEKeysTable(newEKeys)
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

func (parser *ParserModels) UpdateEInfoTableConst(eInfoData []ExchangerInfo) error {
	// Initializing to map to establish faster search from array
	tableDataMap := map[int]ExchangerInfo{}

	// Selected all existing exchanger_info from SQL table,
	// to make sure that we will not INSERT duplicate of any exchanger
	tableData, err := parser.SelectExchangersInfo()
	if err != nil {
		return err
	}

	// tableDataMap is map of tableData -> tableDataMap[id]
	// Writing data from SQL table into map to make search faster
	for _, exchanger := range tableData {
		tableDataMap[exchanger.ExchangerId] = exchanger
	}

	// Finding out all exchangers that was parsed but not found in SQL table -> newEKeys
	var newEInfoData []ExchangerInfo
	for _, exchanger := range eInfoData {
		if _, ok := tableDataMap[exchanger.ExchangerId]; !ok {
			newEInfoData = append(newEInfoData, exchanger)
		}
	}

	// Inserting those new exchanger keys into sql table
	err = parser.InsertEInfoTable(newEInfoData)
	return err
}

func (parser *ParserModels) InsertKZTCurrencies(exchangers []ExchangerCurrencies) error {
	query := fmt.Sprintf(`
		INSERT INTO %s 
		(exchanger_id, upload_time, usd_buy, usd_sell, eur_buy, eur_sell, rub_buy, rub_sell)
		VALUES`,
		postgres.ExchangersCurrenciesTable)

	for i, exchanger := range exchangers {
		query += fmt.Sprintf(`
			(%d,%d,%f,%f,%f,%f,%f,%f)	
			`, exchanger.ExchangerId,
			time.Now().Unix(),
			exchanger.USD_BUY, exchanger.USD_SELL,
			exchanger.EUR_BUY, exchanger.EUR_SELL,
			exchanger.RUB_BUY, exchanger.RUB_SELL)
		if i == len(exchangers)-1 {
			query += ";"
			break
		}
		query += ","
	}

	res, err := parser.psql.Query(query)
	if err != nil {
		return err
	}
	return res.Close()
}

// GetKeysById && GetKeysByName: Methods for interaction with parser,_m1, parser._m2 maps in the higher levels
func (parser *ParserModels) GetKeysById(id int) ExchangerKeys {
	return parser._m1[id]
}

func (parser *ParserModels) GetKeysByName(name string) ExchangerKeys {
	return parser._m2[name]
}

func (parser *ParserModels) InsertEInfoTable(eInfoData []ExchangerInfo) error {
	// Inserting data into exchangers_info table
	if len(eInfoData) == 0 {
		return nil
	}

	query := fmt.Sprintf(`
		INSERT INTO %s 
		(exchanger_id, address, wholesale, updated_time, phone_numbers)
		VALUES`, postgres.ExchangersInfoTable)

	// For loop to multiple insertion on single SLQ request
	for i, eInfo := range eInfoData {
		query += fmt.Sprintf(`
			(%d, '%s', '%s', %d,'%s')`, eInfo.ExchangerId, eInfo.Address, eInfo.WholeSale, eInfo.UpdatedTime, eInfo.PhoneNumbers)

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

func (parser *ParserModels) InsertIntoEKeysTable(eKeys []ExchangerKeys) ([]ExchangerKeys, error) {
	if len(eKeys) == 0 {
		return nil, nil
	}
	var tableEKeys []ExchangerKeys

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

func (parser *ParserModels) SelectExchangersInfo() ([]ExchangerInfo, error) {
	var eInfoData []ExchangerInfo
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ExchangersInfoTable)
	err := parser.psql.Select(&eInfoData, query)
	if err != nil {
		return nil, err
	}
	return eInfoData, nil
}

func (parser *ParserModels) SelectExchangersKeys() ([]ExchangerKeys, error) {
	var eKeys []ExchangerKeys
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ExchangersKeysTable)
	if err := parser.psql.Select(&eKeys, query); err != nil {
		return nil, err
	}
	return eKeys, nil
}

func (parser *ParserModels) SelectExchangersCurrencies() ([]ExchangerCurrencies, error) {
	var data []ExchangerCurrencies
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ExchangersCurrenciesTable)
	if err := parser.psql.Select(&data, query); err != nil {
		return nil, err
	}
	return data, nil
}
