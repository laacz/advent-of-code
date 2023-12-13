package main

import (
	"testing"
)

var INPUT = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := 6032
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := 5031
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
