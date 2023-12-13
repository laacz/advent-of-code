package main

import (
	"testing"
)

var input = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

func TestPartOne(t *testing.T) {
	expect := 1120
	if got := partOne(input, 1000); got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 689
	if got := partTwo(input, 1000); got != expect {
		t.Errorf("PartTwo() = %v, want %v", got, expect)
	}
}
