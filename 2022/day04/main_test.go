package main

import "testing"

var input string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPartOne(t *testing.T) {
	expect := 2
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 4
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
