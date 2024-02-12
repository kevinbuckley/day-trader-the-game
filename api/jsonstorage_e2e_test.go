package main

import (
	"os"
	"testing"
	"time"
	"github.com/joho/godotenv"

)

func Test_SendMarketDataToJSONBin_DataStoredAndBinIdReturned(t *testing.T) {
	err := godotenv.Load()
	apiKey := os.Getenv("JSON_BIN_API_KEY")
	collectionId :=  os.Getenv("JSON_BIN_COLLECTION_ID")
	testData := MarketData{
		AsOf: time.Now(),
		Results: []struct {
			Close  float64
			Ticker string
		}{
			{
				Close:  150.10,
				Ticker: "AAPL",
			},
			{
				Close:  200.20,
				Ticker: "GOOG",
			},
		},
	}
	jsonStorageAPI := BuildDataStorage(apiKey, collectionId)
	binId, err := jsonStorageAPI.SendMarketDataToJSONBin(testData)
	if err != nil { 
		t.Fatalf("%s", err)
	}

	if binId == "" { 
		t.Fatalf("JSONBin is empty")	
	}

}