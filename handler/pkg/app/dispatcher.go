package app

type ExchangerKeys struct {
	Id   int    `json:"id" db:"id"`
	City string `json:"city" db:"city"`
	Name string `json:"name" db:"name"`
}

type ExchangerInfo struct {
	ExchangerId  int    `json:"exchanger_id" db:"exchanger_id"`
	Address      string `json:"address" db:"address"`
	Link         string `json:"link" db:"link"`
	SpecialOffer string `json:"special_offer" db:"special_offer"`
	PhoneNumbers string `json:"phone_numbers" db:"phone_numbers"`
}

type ExchangerCurrencies struct {
	ExchangerId int     `json:"exchanger_id" db:"exchanger_id"`
	USD_BUY     float32 `json:"USD_BUY" db:"usd_buy"`
	USD_SELL    float32 `json:"USD_SELL" db:"usd_sell"`
	EUR_BUY     float32 `json:"EUR_BUY" db:"eur_buy"`
	EUR_SELL    float32 `json:"EUR_SELL" db:"eur_sell"`
	RUB_BUY     float32 `json:"RUB_BUY" db:"rub_buy"`
	RUB_SELL    float32 `json:"RUB_SELL" db:"rub_sell"`
	UpdatedTime int64   `json:"updated_time"`
}
