package main

import "testing"

func TestPartOne(t *testing.T) {
	got := partOne(`ULL
	RRDDD
	LURDL
	UUUUD`)
	if got != "1985" {
		t.Error("Expected 1985, got ", got)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(`ULL
	RRDDD
	LURDL
	UUUUD`)
	if got != "5DB3" {
		t.Error("Expected 5DB3, got ", got)
	}
}
