package models

type ExchangerInfoParams struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Wholesale bool   `json:"wholesale"`
}

type CurrenciesDataParams struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type ExchangerData struct {
	Id           int    `json:"id" db:"id"`
	City         string `json:"city" db:"city"`
	Name         string `json:"name" db:"name"`
	Address      string `json:"address" db:"address"`
	WholeSale    string `json:"wholesale" db:"wholesale"`
	UpdatedTime  uint64 `json:"updated_time" db:"updated_time"`
	PhoneNumbers string `json:"phone_numbers" db:"phone_numbers"`
}

type CurrenciesData struct {
	UploadTime uint64     `json:"upload_time" db:"upload_time"`
	Currencies Currencies `json:"currencies" db:"currencies""`
}

type Currencies struct {
	UploadTime uint64     `json:"upload_time"`
	USD        [2]float32 `json:"USD"`
	EUR        [2]float32 `json:"EUR"`
	RUB        [2]float32 `json:"RUB"`
}

//
//type CurrencyDataParams struct {
//	DataFrom int64 `json:"data_from"`
//	DataTo   int64 `json:"data_to"`
//}

//
//type ExchangersDataByCityResponse struct {
//	Id              int    `json:"id" db:"id"`
//	Name            string `json:"name" db:"name"`
//	CurrencyHistory []Data `json:"currency_history"`
//}
//
