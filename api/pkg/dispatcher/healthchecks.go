package dispatcher

import (
	"api/pkg/models"
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func HealthCheck(dp *NewDispatcher) models.ErrorResponse {
	for i := 0; true; i++ {
		res, err := dp.Get("/healthcheck")
		if i == 10 {
			return models.ErrorResponse{
				Status:  err.Status,
				Message: string(res),
			}
		}

		if err.Message == "" {
			return models.ErrorResponse{}
		}

		time.Sleep(2 * time.Second)
	}

	return models.ErrorResponse{}
}

func (dp *Dispatcher) ServicesShutDownCheck() error {
	var wg sync.WaitGroup
	ch := make(chan error, 1)
	go func() {
		wg.Add(1)
		log.Println("AUTH service shut down started...")
		for i := 0; true; i++ {
			res, err := dp.AuthDispatcher.Get("/healthcheck")
			if i == 10 {
				wg.Done()
				ch <- fmt.Errorf("AUTH service is not shutting down -> Status Code: %d -> Message: %s", err.Status, string(res))
			}
			if err.Message != "" {
				log.Printf(" -> AUTH service shut down. \n")
				wg.Done()
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	if err := <-ch; err != nil {
		return err
	}

	log.Println("HANDLER service shut down started...")
	for i := 0; true; i++ {
		res, err := dp.HandlerDispatcher.Get("/healthcheck")
		if i == 10 {
			return errors.New(fmt.Sprintf("HANDLER service is not shutting down -> Status Code: %d -> Message: %s", err.Status, string(res)))
		}
		if err.Message != "" {
			log.Printf(" -> HANDLER service shut down. \n")
			break
		}
		time.Sleep(time.Second)
	}

	wg.Wait()
	return nil
}
