package main

import "testing"

var input string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

var elves = caloriesPerElf(input)

func TestPartOne(t *testing.T) {
	expect := 24000
	actual := partOne(elves)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 45000
	actual := partTwo(elves)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
