package main

import (
	"strings"
	"testing"
)

var input = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPartOne(t *testing.T) {
	expect := 14
	actual := part1(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 34
	actual := part2(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
