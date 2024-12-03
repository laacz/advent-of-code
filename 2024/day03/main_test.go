package main

import (
	"testing"
)

var input1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
var input2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestPartOne(t *testing.T) {
	expect := 161
	actual := part1(input1)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 48
	actual := part2(input2)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
