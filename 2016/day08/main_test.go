package main

import "testing"

const input = `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4 
rotate column x=1 by 1`

func TestPartOne(t *testing.T) {
	got := partOne(input, 7, 3)
	expected := 6
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(input, 7, 3)
	expected := `.#..#.#
#.#....
.#.....
`
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
