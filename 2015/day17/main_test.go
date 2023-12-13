package main

import (
	"testing"
)

var input = `20
15
10
5
5`

func TestPartOne(t *testing.T) {
	expect := 4
	if got := partOne(input, 25); got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	expect := 3
	if got := partTwo(input, 25); got != expect {
		t.Errorf("PartTwo() = %v, want %v", got, expect)
	}
}
