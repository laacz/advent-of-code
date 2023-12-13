package main

import (
	"testing"
)

var INPUT = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := 3068
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := uint64(1514285714288)
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
