package main

import "testing"

var input string = `A Y
B X
C Z
`

func TestPartOne(t *testing.T) {
	expect := 15
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 12
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
