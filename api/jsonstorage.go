package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
}

type DataStorage struct {
	apiKey              string
	jsonBinCollectionId string
	baseUrl             string
}

func BuildDataStorage(
	apiKey string,
	jsonBinCollectionId string,
	baseUrl string) *DataStorage {
	return &DataStorage{apiKey: apiKey, jsonBinCollectionId: jsonBinCollectionId, baseUrl: baseUrl}
}

func (d *DataStorage) SendMarketDataToJSONBin(data MarketData) (string, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", d.baseUrl+"/v3/b", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Access-Key", d.apiKey)
	req.Header.Set("X-Collection-Id", d.jsonBinCollectionId)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error from JSONBin: %v\n", resp)
		return "", fmt.Errorf("failed to send data to JSONBin, status code: %d", resp.StatusCode)
	} else {
		body, _ := io.ReadAll(resp.Body)
		var result ApiResponse
		if err := json.Unmarshal(body, &result); err != nil {
			return "", err
		}
		return result.Metadata.ID, nil
	}
}
