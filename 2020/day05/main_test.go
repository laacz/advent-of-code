package main

import "testing"

var input = `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

func TestPart1(t *testing.T) {
	expected := 820
	actual := part1(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
