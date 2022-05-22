package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kapralovs/warehouse-client/internal/encoder"
	"github.com/kapralovs/warehouse-client/internal/requests"
)

func main() {
	c := requests.New()
	url := requests.ConstructURL(requests.HOST, requests.PORT, requests.ROUTE)
	encodedCreds := encoder.Encode(os.Args[1], os.Args[2])
	req, err := requests.CreateRequest("GET", url, encodedCreds, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}
