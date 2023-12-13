package main

import (
	"testing"
)

var INPUT = `1
2
-3
3
-2
0
4`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := int64(3)
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := int64(1623178306)
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
