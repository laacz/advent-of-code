package main

import "testing"

var input string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPartOne(t *testing.T) {
	expect := "CMZ"
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := "MCD"
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}
