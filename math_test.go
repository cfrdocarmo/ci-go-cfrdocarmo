package main

import "testing"

// Tested with:
func TestAdd(t *testing.T) {
	if Add(10, 20) != 30 {
		t.Error("Expected 10 + 20 to equal 30")
	}
}
