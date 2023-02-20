package app

type ParserResponse struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	City         string   `json:"city"`
	Address      string   `json:"address"`
	Link         string   `json:"link"`
	SpecialOffer string   `json:"special_offer"`
	PhoneNumbers []string `json:"phone_numbers"`
}
