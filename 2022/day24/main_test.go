package main

import (
	"testing"
)

var INPUT = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := 18
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := 54
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
