package main

import (
	"os"
	"testing"
	"github.com/joho/godotenv"

)

func Test_GetTodaysData_AAPL_ValidTickDataReturned(t *testing.T) {
	err := godotenv.Load()
	apiKey := os.Getenv("POLYGON_API_KEY")
	marketDataAPI := BuildMarketDataAPI(apiKey)
	tickers := []string { "AAPL" }
	tickData, err := marketDataAPI.GetTodaysData(tickers)

	if err != nil { 
		t.Fatalf("%s", err)
	}

	if len(tickData.Results) == 0 || tickData.Results[0].Close <= 0 { 
		t.Fatalf("TickData not valid: %+v", *tickData)	
	}

}