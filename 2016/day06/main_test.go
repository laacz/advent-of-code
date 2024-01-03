package main

import "testing"

const input = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func TestPartOne(t *testing.T) {
	got := partOne(input)
	expected := "easter"
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(input)
	expected := "advent"
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
