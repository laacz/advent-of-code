package main

import (
	"testing"
)

const input = `1
2
3
4
5
7
8
9
10
11`

func TestPartOne(t *testing.T) {
	expect := 99
	got := partOne(input)
	if got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	expect := 44
	got := partTwo(input)
	if got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
