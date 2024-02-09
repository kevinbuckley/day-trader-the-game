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

func (d *DataStorage) SendMarketDataToJSONBin(data MarketData) error {
	jsonData, err := json.Marshal(data)

	fmt.Println(d.apiKey)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Access-Key", d.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error from JSONBin: %v\n", resp)
		return fmt.Errorf("failed to send data to JSONBin, status code: %d", resp.StatusCode)
	}

	fmt.Printf("response from JSONBin: %v\n", resp)

	return nil
}
