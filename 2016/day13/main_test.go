package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `10 7 4`

	expect := 11
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
