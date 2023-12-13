package main

import (
	"testing"
)

var INPUT = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := 33
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(INPUT)
	expect := 3472
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
