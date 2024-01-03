package main

import "testing"

const input = `abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn`

func TestPartOne(t *testing.T) {
	got := partOne(input)
	expected := 2
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPartTwo(t *testing.T) {
	input := `aba[bab]xyz
	xyx[xyx]xyx
	aaa[kek]eke
	zazbz[bzb]cdb`
	got := partTwo(input)
	expected := 3
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
