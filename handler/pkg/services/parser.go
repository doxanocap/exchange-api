package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"handler/pkg/configs"
	"handler/pkg/models"
	"handler/pkg/repository"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// ParserService Works mainly with requesting and handling data from remote service "kzt-parser"
// and writing it into SQL tables
type ParserService struct {
	parserModels repository.Parser
}

func InitParserService(parserModels repository.Parser) *ParserService {
	parseService := &ParserService{parserModels: parserModels}

	exchangers, err1 := parseService.ParseAllExchangers()

	if err1 != nil {
		if err1.Error() != "unmarshal" {
			log.Fatalf("FATAL -> Init parser services -> %s", err1.Error())
		}
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
	ticker2 := time.NewTicker(10 * time.Second)
	err := make(chan error, 4)
	go func() {
		for {
			select {
			case newErr := <- err:
				if newErr != nil {
					log.Printf("Error in attempt to update const data in exchangers info table: %s", newErr.Error())
				}
			case <-ticker1.C:
				data, err1 := parseService.ParseAllExchangers()
				err <- err1 
				err <- parseService.updateEKeysTableConst(data)
				err <- parseService.updateEInfoTableConst(data)				
				log.Printf("!!!IMPORTANT!!! 12 hours passed. Executing table update.")
			case <-ticker2.C:
				err <- parseService.insertKZTCurrencies()
				log.Printf("!!!IMPORTANT!!! Handled currencies data. 10 seconds passed")
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

// ParseAllExchangers Handles parsed data and returns as slice
func (parser *ParserService) ParseAllExchangers() ([]models.ParserResponse, error) {
	var exchangers []models.ParserResponse

	// Hard coding cities, leaving only Astana for a while.
	var cities = []string{"astana"}
	for _, city := range cities {
		res, err := parser.ParseExchangersByCity(city)
		if err != nil {
			return nil, err
		}
		exchangers = append(exchangers, res...)
	}

	return exchangers, nil
}

func (parser *ParserService) ParseExchangersByCity(city string) ([]models.ParserResponse, error) {
	var exchangers []models.ParserResponse
	url := fmt.Sprintf(configs.ENV("DOMAINS_PARSER") + "/kzt-parser/exchangers/" + city)
	// Sending request to separate PARSER service,
	// which parser data about currencies and exchangers from another website
	body, err := parser.getRequest(url)
	if err != nil {
		log.Println("qweqweq:",err)
		return nil, err
	}

	// Handling json response body

	if err := json.Unmarshal(body, &exchangers); err != nil {
		log.Printf("Unable to unmarshal following %s", err.Error())
		return nil, errors.New("unmarshal")
	}
	return exchangers, nil
}

func (parser *ParserService) GetExchangersKeysTable() ([]repository.ExchangerKeys, error) {
	response, err := parser.parserModels.SelectExchangersKeys()
	if err != nil {
		return nil, err
	}
	return response, err
}

func (parser *ParserService) GetExchangersInfoTable() ([]repository.ExchangerInfo, error) {
	response, err := parser.parserModels.SelectExchangersInfo()
	if err != nil {
		return nil, err
	}
	return response, err
}

func (parser *ParserService) GetExchangersCurrenciesTable() ([]models.ExchangerCurrenciesResponse, error) {
	data, err := parser.parserModels.SelectExchangersCurrencies()
	if err != nil {
		return nil, err
	}

	var response []models.ExchangerCurrenciesResponse
	for _, e := range data {
		response = append(response, models.ExchangerCurrenciesResponse{
			ExchangerId: e.ExchangerId,
			UploadTime:  e.UploadTime,
			USD:         [2]float32{e.USD_BUY, e.USD_SELL},
			EUR:         [2]float32{e.EUR_BUY, e.EUR_SELL},
			RUB:         [2]float32{e.RUB_BUY, e.RUB_SELL},
		})
	}
	return response, err
}

//
//
// @Bellow are only private methods for this service
//
//

func (parser *ParserService) insertKZTCurrencies() error {
	var eCurrencies []repository.ExchangerCurrencies
	exchangers, err := parser.ParseAllExchangers()
	for _, exchanger := range exchangers {
		eCurrencies = append(eCurrencies, repository.ExchangerCurrencies{
			parser.parserModels.GetKeysByName(exchanger.Name).Id,
			exchanger.UpdatedTime,
			exchanger.USD[0], exchanger.USD[1],
			exchanger.EUR[0], exchanger.EUR[1],
			exchanger.RUB[0], exchanger.RUB[1],
		})
	}

	if err != nil {
		return err
	}
	err = parser.parserModels.InsertKZTCurrencies(eCurrencies)
	return err
}

func (parser *ParserService) updateEInfoTableConst(exchangers []models.ParserResponse) error {
	var eInfo []repository.ExchangerInfo
	for _, exchanger := range exchangers {
		info := repository.ExchangerInfo{
			ExchangerId:  parser.parserModels.GetKeysByName(exchanger.Name).Id,
			Address:      exchanger.Address,
			WholeSale:    exchanger.WholeSale,
			PhoneNumbers: strings.Join(exchanger.PhoneNumbers, ","),
		}
		eInfo = append(eInfo, info)
	}

	err := parser.parserModels.UpdateEInfoTableConst(eInfo)
	return err
}

func (parser *ParserService) updateEKeysTableConst(exchangers []models.ParserResponse) error {
	var eKeys []repository.ExchangerKeys
	for _, exchanger := range exchangers {
		eKey := repository.ExchangerKeys{
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
