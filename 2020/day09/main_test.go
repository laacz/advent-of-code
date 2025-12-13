package main

import "testing"

var input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestPart1(t *testing.T) {
	expected := 127
	actual := part1(parse(input), 5)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 62
	actual := part2(parse(input), 5)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
