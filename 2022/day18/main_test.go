package main

import (
	"testing"
)

var INPUT = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := 64
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := 58
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
