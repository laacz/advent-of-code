package main

import (
	"testing"
)

//var INPUT = `.....
//..##.
//..#..
//.....
//..##.
//.....`

var INPUT = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := 110
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := 20
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
