package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

	expect := 14
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
