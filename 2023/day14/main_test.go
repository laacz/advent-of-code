package main

import (
	"testing"
)

var input = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestPartOne(t *testing.T) {
	expected := 136
	actual := partOne(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 64
	actual := partTwo(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
