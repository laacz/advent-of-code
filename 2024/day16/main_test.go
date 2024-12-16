package main

import (
	"testing"
)

var tests = []struct {
	input string
	part1 int
	part2 int
}{
	{
		`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
		7036, 45,
	},
	{
		`#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`,
		11048, 64,
	},
}

func TestPartOne(t *testing.T) {
	for _, test := range tests {
		maze, sart, exit := parseInput(test.input)
		actual := part1(maze, sart, exit)
		if actual != test.part1 {
			t.Errorf("Expected %d, got %d", test.part1, actual)
			// break
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, test := range tests {
		maze, sart, exit := parseInput(test.input)
		actual := part2(maze, sart, exit)
		if actual != test.part2 {
			t.Errorf("Expected %d, got %d", test.part2, actual)
		}
	}
}
