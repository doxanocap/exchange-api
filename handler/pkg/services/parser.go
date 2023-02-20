package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"handler/pkg/app"
	"handler/pkg/repository"
	"io/ioutil"
	"net/http"
	"time"
)

// ParserService Works mainly with requesting and handling data from remote service "kzt-parser"
// and writing it into SQL tables
type ParserService struct {
	parserModels repository.ParserModels
}

func InitParserService(parserModels *repository.ParserModels) *ParserService {
	parseService := &ParserService{parserModels: *parserModels}

	// Initializing parser exchangers_keys tables data.
	// If something wrong during init process, server shut down immediately.
	if err := parseService.updateExchangersTableConst(); err != nil {
		logrus.Fatalf("Error in init parser services: %s", err.Error())
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
					logrus.Printf("Error in attempt to update const data in exchangers table: %s", errMsg.Error())
				}
			case <-ticker1.C:
				err <- parseService.updateExchangersTableConst()
				logrus.Printf("!!!IMPORTANT!!! 12 hours passed. Executing table update.")
			case <-ticker2.C:
				err <- parseService.InsertKZTCurrencies()
				logrus.Printf("!!!IMPORTANT!!! Handled currencies data. 5 minutes passed")
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
	exchangers, err := parser.GetAllExchangers()
	if err != nil {
		return err
	}
	err = parser.parserModels.InsertKZTCurrencies(exchangers)
	return err
}

// GetAllExchangers Handles parsed data and returns as slice
func (parser *ParserService) GetAllExchangers() ([]app.Exchanger, error) {
	var exchangers []app.Exchanger

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

func (parser *ParserService) GetExchangersByCity(city string) ([]app.Exchanger, error) {
	var exchangers []app.Exchanger

	url := fmt.Sprintf(viper.GetString("domains.services.kzt-parser") + "/" + city)

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

func (parser *ParserService) updateExchangersTableConst() error {
	exchangers, err := parser.GetAllExchangers()
	if err != nil {
		return err
	}

	err = parser.parserModels.UpdateExchangersTableConst(exchangers)
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
