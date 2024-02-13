package main

import (
	"testing"
)

// equals fails the test if got is not equal to want.
func equals(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
