package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// import necessary packages

type MarketDataResponse struct {
	Results []struct {
		Close float64 `json:"c"`
	} `json:"results"`
}

type MarketDataAPI struct {
	apiKey string
}

func BuildMarketDataAPI(apiKey string) *MarketDataAPI {
	return &MarketDataAPI{apiKey: apiKey}
}

func (m *MarketDataAPI) GetEndOfDay(ticker string) (*MarketDataResponse, error) {
	url := fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/%s/prev?adjusted=true&apiKey=%s", ticker, m.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response MarketDataResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
