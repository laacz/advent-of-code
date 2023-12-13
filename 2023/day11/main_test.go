package main

import (
	"testing"
)

var input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPartOne(t *testing.T) {
	expect := 374
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 82000210
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
