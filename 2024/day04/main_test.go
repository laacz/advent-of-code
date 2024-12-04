package main

import (
	"testing"
)

var input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPartOne(t *testing.T) {
	expect := 18
	actual := part1(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 9
	actual := part2(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
