package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	tickers := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "FB"}
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	apiKey := os.Getenv("POLYGON_API_KEY")
	if apiKey == "" {
		panic("POLYGON_API_KEY environment variable not set")
	}
	marketDataAPI := BuildMarketDataAPI(apiKey)
	todaysData, err := marketDataAPI.GetTodaysData(tickers)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", todaysData)
}
