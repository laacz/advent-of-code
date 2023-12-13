package main

import (
	"testing"
)

var input string = `30373
25512
65332
33549
35390`

func TestPartOne(t *testing.T) {
	got := partOne(input)
	expect := 21
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(input)
	expect := 8
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
