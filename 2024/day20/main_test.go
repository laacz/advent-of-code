package main

import (
	"testing"
)

var input = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestPartOne(t *testing.T) {
	comp := parseInput(input)
	expected := 44
	if actual := part1(comp, 3); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	comp := parseInput(input)
	expected := 285
	if actual := part2(comp, 50); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
