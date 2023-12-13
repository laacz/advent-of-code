package main

import (
	"testing"
)

var input string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPartOne(t *testing.T) {
	expect := 114
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestSumToZero(t *testing.T) {
	expect := 590376
	actual := partOne("6 19 46 99 204 412 817 1580 2958 5337 9268 15505 25044 39162 59455 87874 126758 178863 247386 335983 448780")

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 2
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
