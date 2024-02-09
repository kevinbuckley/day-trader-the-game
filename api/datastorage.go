package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DataStorage struct {
	apiKey string
}

func BuildDataStorage(apiKey string) *DataStorage {
	return &DataStorage{apiKey: apiKey}
}

func (d *DataStorage) sendMarketDataToJSONBin(data MarketData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", d.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send data to JSONBin, status code: %d", resp.StatusCode)
	}

	return nil
}
