package main

import (
	"testing"
)

func TestCalculateFrequency(t *testing.T) {
	expected := 561

	res := calculate("input.txt")
	if expected != res {
		t.Errorf("Expected %d got %d", expected, res)
	}
}
