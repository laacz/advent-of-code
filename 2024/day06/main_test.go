package main

import (
	"strings"
	"testing"
)

var input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPartOne(t *testing.T) {
	expect := 41
	actual := part1(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 6
	actual := part2(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
