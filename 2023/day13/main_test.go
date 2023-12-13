package main

import (
	"testing"
)

var input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestPartOne(t *testing.T) {
	expected := 405
	actual := partOne(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 400
	actual := partTwo(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
