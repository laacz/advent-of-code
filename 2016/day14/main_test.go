package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `abc`

	expect := 22728
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := `abc`

	expect := 22551
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
