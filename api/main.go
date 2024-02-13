package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getAndStoreMarketData(m *MarketDataAPI, d *DataStorage, tickers []string) string {
	todaysData, err := m.GetTodaysData(tickers)
	if err != nil {
		panic(err)
	}
	binId, err := d.SendMarketDataToJSONBin(*todaysData)
	if err != nil {
		panic(err)
	}
	return binId
}

func getRequiredEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("%s environment variable not set", key))
	}
	return value
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	polygonKey := getRequiredEnvVar("POLYGON_API_KEY")
	jsonBinKey := getRequiredEnvVar("JSON_BIN_API_KEY")
	jsonBinCollectionId := getRequiredEnvVar("JSON_BIN_COLLECTION_ID")

	tickers := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "FB"}
	marketDataAPI := BuildMarketDataAPI(polygonKey, "https://api.polygon.io")
	dataStorage := BuildDataStorage(jsonBinKey, jsonBinCollectionId, "https://api.jsonbin.io")
	binId := getAndStoreMarketData(marketDataAPI, dataStorage, tickers)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Market data stored in JSONBin with ID: %s\n", binId)
	}
}
