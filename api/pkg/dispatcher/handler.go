package dispatcher

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/cookiejar"

	"time"
)

func NewHandlerDispatcher(url string) *NewDispatcher {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	handler := &NewDispatcher{
		url: url,
		client: http.Client{
			Jar:     jar,
			Timeout: time.Second * 10,
		},
	}

	log.Println("handler service healthcheck started...")
	{
		if err := HealthCheck(handler); err.IsError() {
			log.Fatalf("handler service is not shutting down -> Status Code: %d -> Message: %s", err.Status, err.Message)
		}
	}
	return handler
}
