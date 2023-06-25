package models

type ExchangersResponse struct {
	Id           int    `json:"id" db:"id"`
	City         string `json:"city" db:"city"`
	Name         string `json:"name" db:"name"`
	Address      string `json:"address" db:"address"`
	WholeSale    string `json:"wholesale" db:"wholesale"`
	UpdatedTime  uint64 `json:"updated_time" db:"updated_time"`
	PhoneNumbers string `json:"phone_numbers" db:"phone_numbers"`
}

type CurrenciesResponse struct {
	UploadTime uint64     `json:"upload_time" db:"upload_time"`
	Currencies Currencies `json:"currencies" db:"currencies""`
}

type Currencies struct {
	UploadTime uint64     `json:"-"`
	USD        [2]float32 `json:"USD"`
	EUR        [2]float32 `json:"EUR"`
	RUB        [2]float32 `json:"RUB"`
}
