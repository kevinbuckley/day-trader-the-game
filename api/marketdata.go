package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MarketDataResponse struct {
	Results []struct {
		Close float64 `json:"c"`
	} `json:"results"`
}

type MarketData struct {
	AsOf    time.Time
	Results []struct {
		Close  float64
		Ticker string
	}
}

type MarketDataAPI struct {
	apiKey string
}

func BuildMarketDataAPI(apiKey string) *MarketDataAPI {
	return &MarketDataAPI{apiKey: apiKey}
}

func (m *MarketDataAPI) GetTodaysData(tickers []string) (*MarketData, error) {
	var marketData MarketData
	marketData.AsOf = time.Now()
	for _, ticker := range tickers {
		singleTick, err := m.getEndOfDay(ticker)
		if err != nil {
			return nil, err
		}
		if len(singleTick.Results) > 0 {
			marketData.Results = append(marketData.Results, struct {
				Close  float64
				Ticker string
			}{
				Close:  singleTick.Results[0].Close,
				Ticker: ticker,
			})
		}
	}
	return &marketData, nil
}

func (m *MarketDataAPI) getEndOfDay(ticker string) (*MarketDataResponse, error) {
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
