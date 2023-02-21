package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"handler/pkg/app"
	"handler/pkg/repository"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ParserService Works mainly with requesting and handling data from remote service "kzt-parser"
// and writing it into SQL tables
type ParserService struct {
	parserModels repository.ParserModels
}

func InitParserService(parserModels *repository.ParserModels) *ParserService {
	parseService := &ParserService{parserModels: *parserModels}

	exchangers, err1 := parseService.GetAllExchangers()
	if err1 != nil {
		log.Fatalf("FATAL -> Init parser services -> %s", err1.Error())
	}

	// Initializing parser exchangers_keys tables data.
	// If something wrong during init process, server shut down immediately.
	{
		if err := parseService.updateEKeysTableConst(exchangers); err != nil {
			log.Fatalf("FATAL -> Init exchanger_keys table services -> %s", err.Error())
		}
		if err := parseService.updateEInfoTableConst(exchangers); err != nil {
			log.Fatalf("FATAL -> Init exchangers_info table services -> %s", err.Error())
		}
	}

	// Updates data about exchangers every 12 hours.
	ticker1 := time.NewTicker(12 * time.Hour)
	// Parses currencies every 5 minutes and handles data into SQL tables
	ticker2 := time.NewTicker(5 * time.Minute)

	err := make(chan error)
	go func() {
		for {
			select {
			case errMsg := <-err:
				if errMsg != nil {
					log.Printf("Error in attempt to update const data in exchangers table: %s", errMsg)
				}
			case <-ticker1.C:
				err <- parseService.updateEKeysTableConst(exchangers)
				err <- parseService.updateEInfoTableConst(exchangers)
				log.Printf("!!!IMPORTANT!!! 12 hours passed. Executing table update.")
			case <-ticker2.C:
				err <- parseService.InsertKZTCurrencies()
				log.Printf("!!!IMPORTANT!!! Handled currencies data. 5 minutes passed")
			}
		}
	}()

	return parseService
}

//
//
// @Bellow are only public methods for this service
//
//

func (parser *ParserService) InsertKZTCurrencies() error {
	var eCurrencies []app.ExchangerCurrencies
	exchangers, err := parser.GetAllExchangers()
	for _, exchanger := range exchangers {
		eCurrencies = append(eCurrencies, app.ExchangerCurrencies{
			parser.parserModels.GetKeysByName(exchanger.Name).Id,
			exchanger.USD_BUY, exchanger.USD_SELL,
			exchanger.EUR_BUY, exchanger.EUR_SELL,
			exchanger.RUB_BUY, exchanger.RUB_SELL,
			time.Now().Unix(),
		})
	}
	if err != nil {
		return err
	}
	err = parser.parserModels.InsertKZTCurrencies(eCurrencies)
	return err
}

// GetAllExchangers Handles parsed data and returns as slice
func (parser *ParserService) GetAllExchangers() ([]app.ParserResponse, error) {
	var exchangers []app.ParserResponse

	// Hard coding cities, leaving only Astana for a while.
	var cities = []string{"astana"}
	for _, city := range cities {
		res, err := parser.GetExchangersByCity(city)
		if err != nil {
			return nil, err
		}
		exchangers = append(exchangers, res...)
	}

	return exchangers, nil
}

func (parser *ParserService) GetExchangersByCity(city string) ([]app.ParserResponse, error) {
	var exchangers []app.ParserResponse
	url := fmt.Sprintf(app.Public.ENV["DOMAINS.KZT_PARSER"] + "/exchangers/" + city)
	// Sending request to separate PARSER service,
	// which parser data about currencies and exchangers from another website
	body, err := parser.getRequest(url)
	if err != nil {
		return nil, err
	}

	// Handling json response body
	if err := json.Unmarshal(body, &exchangers); err != nil {
		return nil, err
	}
	return exchangers, nil
}

//
//
// @Bellow are only private methods for this service
//
//

func (parser *ParserService) updateEInfoTableConst(exchangers []app.ParserResponse) error {
	var eInfo []app.ExchangerInfo
	for _, exchanger := range exchangers {
		info := app.ExchangerInfo{
			ExchangerId:  parser.parserModels.GetKeysByName(exchanger.Name).Id,
			Address:      exchanger.Address,
			Link:         exchanger.Link,
			SpecialOffer: exchanger.SpecialOffer,
			PhoneNumbers: strings.Join(exchanger.PhoneNumbers, ","),
		}
		eInfo = append(eInfo, info)
	}

	err := parser.parserModels.UpdateEInfoTableConst(eInfo)
	return err
}

func (parser *ParserService) updateEKeysTableConst(exchangers []app.ParserResponse) error {
	var eKeys []app.ExchangerKeys
	for _, exchanger := range exchangers {
		eKey := app.ExchangerKeys{
			Id:   0,
			City: exchanger.City,
			Name: exchanger.Name,
		}
		eKeys = append(eKeys, eKey)
	}

	err := parser.parserModels.UpdateEKeysTableConst(eKeys)
	return err
}

// Method for sending GET request
func (parser *ParserService) getRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Method for sending POST request
func (parser *ParserService) postRequest(url string, postBody []byte) ([]byte, error) {
	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBufferString(string(postBody)))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
