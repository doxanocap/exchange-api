package dispatcher

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func (dp *Dispatcher) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "sending get request")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response from get")
	}

	return body, nil
}

// Post Method for sending POST request
func (dp *Dispatcher) Post(url string, postBody []byte) ([]byte, error) {
	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBufferString(string(postBody)))

	if err != nil {
		return nil, errors.Wrap(err, "sending post request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response from post")
	}

	return body, nil
}
