package app

type Exchanger struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	City         string   `json:"city"`
	Address      string   `json:"address"`
	Link         string   `json:"link"`
	SpecialOffer string   `json:"special_offer"`
	PhoneNumbers []string `json:"phone_numbers"`
	USD_BUY      float32  `json:"usd_buy"`
	USD_SELL     float32  `json:"usd_sell"`
	EUR_BUY      float32  `json:"eur_buy"`
	EUR_SELL     float32  `json:"eur_sell"`
	RUB_BUY      float32  `json:"rub_buy"`
	RUB_SELL     float32  `json:"rub_sell"`
	UpdatedTime  int64    `json:"updated_time"`
}

type ExchangerKeys struct {
	Id   int    `json:"id"`
	City string `json:"city"`
	Name string `json:"name"`
}

type ExchangerInfo struct {
	Id           int      `json:"id"`
	Address      string   `json:"address"`
	Link         string   `json:"link"`
	SpecialOffer string   `json:"special_offer"`
	PhoneNumbers []string `json:"phone_numbers"`
}

type ExchangerCurrencies struct {
	ExchangerId int     `json:"exchanger_id"`
	USD_BUY     float32 `json:"usd_buy"`
	USD_SELL    float32 `json:"usd_sell"`
	EUR_BUY     float32 `json:"eur_buy"`
	EUR_SELL    float32 `json:"eur_sell"`
	RUB_BUY     float32 `json:"rub_buy"`
	RUB_SELL    float32 `json:"rub_sell"`
	UpdatedTime int64   `json:"updated_time"`
}
