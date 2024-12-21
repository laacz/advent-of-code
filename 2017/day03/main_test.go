package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, test := range []struct {
		input  int
		expect int
	}{
		{1, 0},
		{2, 1},
		{12, 3},
		{15, 2},
		{21, 4},
		{23, 2},
		{25, 4},
		{26, 5},
		{1024, 31},
	} {
		if actual := part1(test.input); actual != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, test := range []struct {
		input  int
		expect int
	}{
		{1, 2},
		{2, 4},
		{12, 23},
		{15, 23},
		{21, 23},
		{23, 25},
		{25, 26},
		{26, 54},
		{1024, 1968},
	} {
		if actual := part2(test.input); actual != test.expect {
			t.Errorf("For %d expected %d, got %d", test.input, test.expect, actual)
		}
	}
}

func TestSpiralCoord(t *testing.T) {
	for _, test := range []struct {
		input  int
		expect Coord
	}{
		{1, Coord{0, 0}},
		{2, Coord{1, 0}},
		{3, Coord{1, -1}},
		{4, Coord{0, -1}},
		{5, Coord{-1, -1}},
		{24, Coord{1, 2}},
		{25, Coord{2, 2}},
		{18, Coord{-2, -1}},
	} {
		if actual := SpiralCoord(test.input); actual != test.expect {
			t.Errorf("For %d expected %v, got %v", test.input, test.expect, actual)
		}
	}
}
