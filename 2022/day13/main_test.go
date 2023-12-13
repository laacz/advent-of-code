package main

import (
	"testing"
)

var INPUT = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestPartOne(t *testing.T) {
	input := INPUT
	got := partOne(input)
	expect := 13
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	input := INPUT
	got := partTwo(input)
	expect := 140
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
