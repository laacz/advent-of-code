package main

import "testing"

func TestPartOne(t *testing.T) {
	got := partOne(`5 10 25`)
	if got != 0 {
		t.Error("Expected 0, got ", got)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(`101 301 501
	102 302 502
	103 303 503
	201 401 601
	202 402 602
	203 403 603`)
	if got != 0 {
		t.Error("Expected 5DB3, got ", got)
	}
}
