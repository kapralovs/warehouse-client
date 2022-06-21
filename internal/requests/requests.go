package requests

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func New() http.Client {
	client := http.Client{}
	return client
}

func CreateRequest(method, url string, data interface{}) (*http.Request, error) {
	if data != nil || method == "POST" {
		dataAsBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, url, strings.NewReader(string(dataAsBytes)))
		if err != nil {
			return nil, err
		}

		// req.Header.Add("Authorization", "Basic "+encodedCreds)
		req.SetBasicAuth(os.Args[1], os.Args[2])
		req.Header.Add("Content-Type", "application/json")
		return req, nil
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Args[1], os.Args[2])

	return req, nil
}
