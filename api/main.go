package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getAndStoreMarketData(m *MarketDataAPI, d *DataStorage, tickers []string) error {
	todaysData, err := m.GetTodaysData(tickers)
	if err != nil {
		return err
	}
	fmt.Printf("Retrieved market data: \n\n %v\n", todaysData)
	err = d.sendMarketDataToJSONBin(*todaysData)
	if err != nil {
		return err
	}
	fmt.Printf("Retrieved and stored market data: \n\n %v\n", todaysData)
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	polygonKey := os.Getenv("POLYGON_API_KEY")
	if polygonKey == "" {
		panic("POLYGON_API_KEY environment variable not set")
	}
	jsonBinKey := os.Getenv("JSON_BIN_API_KEY")
	if jsonBinKey == "" {
		panic("JSON_BIN_API_KEY environment variable not set")
	}

	tickers := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "FB"}
	marketDataAPI := BuildMarketDataAPI(polygonKey)
	dataStorage := BuildDataStorage(jsonBinKey)
	getAndStoreMarketData(marketDataAPI, dataStorage, tickers)
}
