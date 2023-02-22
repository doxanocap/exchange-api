package models

type ParserResponse struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	City         string   `json:"city"`
	Address      string   `json:"address"`
	Link         string   `json:"link"`
	SpecialOffer string   `json:"special_offer"`
	UpdateTime   string   `json:"update_time"`
	USD_BUY      float32  `json:"USD_BUY"`
	USD_SELL     float32  `json:"USD_SELL"`
	EUR_BUY      float32  `json:"EUR_BUY"`
	EUR_SELL     float32  `json:"EUR_SELL"`
	RUB_BUY      float32  `json:"RUB_BUY"`
	RUB_SELL     float32  `json:"RUB_SELL"`
	PhoneNumbers []string `json:"phone_numbers"`
	UpdatedTime  int64    `json:"updated_time"`
}
