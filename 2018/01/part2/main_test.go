package main

import (
	"testing"
)

func TestCalculateReachesTwiceFrequency(t *testing.T) {
	expected := 563

	res := calculate("input.txt")
	if expected != res {
		t.Errorf("Expected %d got %d", expected, res)
	}
}
