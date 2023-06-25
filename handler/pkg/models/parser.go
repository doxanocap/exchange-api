package models

type ParserResponse struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	City         string     `json:"city"`
	Address      string     `json:"address"`
	WholeSale    string     `json:"wholesale"`
	UpdatedTime  uint64     `json:"updated_time"`
	PhoneNumbers []string   `json:"phone_numbers"`
	USD          [2]float32 `json:"USD"`
	RUB          [2]float32 `json:"RUB"`
	EUR          [2]float32 `json:"EUR"`
}

type ExchangerCurrenciesResponse struct {
	ExchangerId int        `json:"-" db:"exchanger_id"`
	UploadTime  uint64     `json:"upload_time" db:"upload_time"`
	USD         [2]float32 `json:"USD" db:"usd"`
	EUR         [2]float32 `json:"EUR" db:"eur"`
	RUB         [2]float32 `json:"RUB" db:"rub"`
}
