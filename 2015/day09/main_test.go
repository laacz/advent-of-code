package main

import (
	"testing"
)

var input = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141"`

func TestPartOne(t *testing.T) {
	expected := 605
	actual := partOne(input)

	if actual != expected {
		t.Errorf("PartOne(%s) = %v, expected %v", input, actual, expected)
	}

}

func TestPartTwo(t *testing.T) {
	expected := 982
	actual := partTwo(input)
	if actual != expected {
		t.Errorf("PartTwo(%s) = %v, expected %v", input, actual, expected)
	}
}
