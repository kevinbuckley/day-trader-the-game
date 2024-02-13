package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMarketDataToJSONBin(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/v3/b")
		equals(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"metadata": {"id": "test_id"}}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use server.URL as the base URL for your API
	dataStorage := BuildDataStorage("TEST_KEY", "TEST_COLLECTION_ID", server.URL)
	resp, err := dataStorage.SendMarketDataToJSONBin(MarketData{})
	if err != nil {
		t.Fatal(err)
	}

	// Test response data
	equals(t, resp, "test_id")
}
