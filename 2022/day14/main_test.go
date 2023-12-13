package main

import (
	"testing"
)

var INPUT = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestPartOne(t *testing.T) {
	input := INPUT
	got := partOne(input)
	expect := 24
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	input := INPUT
	got := partTwo(input)
	expect := 93
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
