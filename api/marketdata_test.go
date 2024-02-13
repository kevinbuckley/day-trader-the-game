package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEndOfDay(t *testing.T) {
	apiKey := "TEST_KEY"
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/v2/aggs/ticker/TEST/prev?adjusted=true&apiKey="+apiKey)
		// Send response to be tested
		rw.Write([]byte(`{"status": "OK", "results": [{"c": 100.0}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use server.URL as the base URL for your API
	api := BuildMarketDataAPI(apiKey, server.URL)
	resp, err := api.getEndOfDay("TEST")
	if err != nil {
		t.Fatal(err)
	}

	equals(t, resp.Results[0].Close, 100.0)
}
