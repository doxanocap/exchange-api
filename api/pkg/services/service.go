package services

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func SendPostRequest(url string, postBody []byte) []byte {

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBufferString(string(postBody)))

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func SendGetRequest(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
