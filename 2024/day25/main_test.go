package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{
			`#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`,
			3,
		},
	} {
		locks, keys := parseInput(test.input)
		if actual := part1(locks, keys); actual != test.expected {
			t.Errorf("expected %v but got %v", test.expected, actual)
		}
	}
}
