package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `Disc #1 has 5 positions; at time=0, it is at position 4.
	Disc #2 has 2 positions; at time=0, it is at position 1.`

	expect := 5
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := `Disc #1 has 5 positions; at time=0, it is at position 4.
	Disc #2 has 2 positions; at time=0, it is at position 1.`

	expect := 85
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
