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
	marketData := BuildMarketDataAPI(apiKey)
	for _, ticker := range tickers {
		response, err := marketData.GetEndOfDay(ticker)
		if err != nil {
			panic(err)
		}
		for _, result := range response.Results {
			fmt.Printf("Ticker: %s, Close: %f\n", ticker, result.Close)
		}
	}
}
