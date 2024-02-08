package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Results []struct {
		Close float64 `json:"c"`
	} `json:"results"`
}

func queryAPI(ticker, apiKey string) (*Response, error) {
	url := fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/%s/prev?adjusted=true&apiKey=%s", ticker, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func main() {
	var tickers = []string{"AAPL", "MSFT", "GOOGL", "AMZN", "FB"}

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	apiKey := os.Getenv("POLYGON_API_KEY")
	if apiKey == "" {
		panic("POLYGON_API_KEY environment variable not set")
	}
	for _, ticker := range tickers {
		response, err := queryAPI(ticker, apiKey)
		if err != nil {
			panic(err)
		}

		for _, result := range response.Results {
			fmt.Printf("Ticker: %s, Close: %f\n", ticker, result.Close)
		}
	}
}
