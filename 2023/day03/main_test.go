package main

import (
	"testing"
)

var input string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestPartOne(t *testing.T) {
	expect := 4361
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 467835
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
