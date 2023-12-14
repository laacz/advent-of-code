package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expect := 6
	if got := partOne(`120`); got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	expect := 6
	if got := partTwo(`e => H
	e => O
	H => HO
	H => OH
	O => HH
	
	HOHOHO`); got != expect {
		t.Errorf("PartTwo() = %v, want %v", got, expect)
	}
}
