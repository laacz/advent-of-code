package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `value 5 goes to bot 2
	bot 2 gives low to bot 1 and high to bot 0
	value 3 goes to bot 1
	bot 1 gives low to output 1 and high to bot 0
	bot 0 gives low to output 2 and high to output 0
	value 2 goes to bot 2`

	expect := 2
	actual := partOne(input, 2, 5)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := `value 5 goes to bot 2
	bot 2 gives low to bot 1 and high to bot 0
	value 3 goes to bot 1
	bot 1 gives low to output 1 and high to bot 0
	bot 0 gives low to output 2 and high to output 0
	value 2 goes to bot 2`

	expect := 2 * 3 * 5
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
