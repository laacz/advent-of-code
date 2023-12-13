package main

import (
	"testing"
)

var input = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

func TestPartOne(t *testing.T) {
	expect := 62842880
	if got := partOne(input); got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	expect := 57600000
	if got := partTwo(input); got != expect {
		t.Errorf("PartTwo() = %v, want %v", got, expect)
	}
}
