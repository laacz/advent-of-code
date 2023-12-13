package main

import (
	"testing"
)

var input = `""
"abc"
"aaa\"aaa"
"\x27"`

func TestPartOne(t *testing.T) {
	expected := 12
	actual := partOne(input)

	if actual != expected {
		t.Errorf("PartOne(%s) = %v, expected %v", input, actual, expected)
	}

}

func TestPartTwo(t *testing.T) {
	expected := 19
	actual := partTwo(input)
	if actual != expected {
		t.Errorf("PartTwo(%s) = %v, expected %v", input, actual, expected)
	}
}
