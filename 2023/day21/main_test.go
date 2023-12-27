package main

import (
	"testing"
)

var input = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestPartOne(t *testing.T) {
	actual := partOne(input, 6)
	expected := 16

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	var cases = []struct {
		steps    int
		expected int
	}{
		{6, 16},
		{10, 50},
		{50, 1594},
		{100, 6536},
		{500, 167004},
		{1000, 668697},
		{5000, 16733044},
	}

	for _, c := range cases {
		actual := partTwo(input, c.steps)
		if actual != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, actual)
		}
	}
}
