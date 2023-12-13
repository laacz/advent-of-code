package main

import (
	"testing"
)

var input = `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

func TestPartOne(t *testing.T) {
	expect := 4
	if got := partOne(input, 4); got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	expect := 17
	if got := partTwo(input, 5); got != expect {
		t.Errorf("PartTwo() = %v, want %v", got, expect)
	}
}
