package main

import "testing"

func TestPartOne(t *testing.T) {
	got := partOne(`abc`)
	expected := "18f47a30"
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(`abc`)
	expected := "05ace8e3"
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
