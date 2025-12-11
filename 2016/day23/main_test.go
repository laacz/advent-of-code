package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `cpy 2 a
	tgl a
	tgl a
	tgl a
	cpy 1 a
	dec a
	dec a`

	expect := 3
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
