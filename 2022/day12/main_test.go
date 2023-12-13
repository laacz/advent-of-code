package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	got := partOne(input)
	expect := 31
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	got := partTwo(input)
	expect := 29
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
