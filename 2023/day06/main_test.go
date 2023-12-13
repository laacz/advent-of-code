package main

import (
	"testing"
)

var input string = `Time:      7  15   30
Distance:  9  40  200`

func TestPartOne(t *testing.T) {
	expect := 288
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 71503
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
