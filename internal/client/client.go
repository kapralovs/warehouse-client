package client

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func New() http.Client {
	client := http.Client{}
	return client
}

func (c *http.Client) MakeRequest(method, url, encodedCreds string, data interface{}) (*http.Response, error) {
	if data != nil || method == "POST" {
		dataAsBytes, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}

		req, err := http.NewRequest(method, url, strings.NewReader(string(dataAsBytes)))
		if err != nil {
			log.Println(err)
		}

		req.Header.Add("Authorization", "Basic "+encodedCreds)
		req.Header.Add("Content-Type", "application/json")

		return c.Do(req)
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+encodedCreds)

	return c.Do(req)
}
