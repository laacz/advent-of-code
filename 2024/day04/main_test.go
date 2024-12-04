package main

import (
	"testing"
)

var input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPartOne(t *testing.T) {
	expect := 18
	actual := part1(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 9
	actual := part2(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestHasWord(t *testing.T) {
	m, _ := parseInput(input)
	tests := []struct {
		pos  Coord
		dir  Coord
		word string
	}{
		{Coord{9, 9}, directions["N"], "XMAS"},
		{Coord{5, 0}, directions["E"], "XMAS"},
		{Coord{9, 3}, directions["S"], "XMAS"},
		{Coord{4, 1}, directions["W"], "XMAS"},
		{Coord{1, 9}, directions["NE"], "XMAS"},
		{Coord{4, 0}, directions["SE"], "XMAS"},
		{Coord{9, 3}, directions["SW"], "XMAS"},
		{Coord{9, 9}, directions["NW"], "XMAS"},
	}

	for _, test := range tests {
		actual := m.HasWord(test.pos, test.dir, test.word)
		if actual != 1 {
			t.Errorf("Expected 1, got %d", actual)
		}
	}
}

func TestGetWord(t *testing.T) {
	m, _ := parseInput(input)
	tests := []struct {
		pos  Coord
		dir  Coord
		word string
	}{
		{Coord{9, 9}, directions["N"], "XMAS"},
		{Coord{5, 0}, directions["E"], "XMAS"},
		{Coord{9, 3}, directions["S"], "XMAS"},
		{Coord{4, 1}, directions["W"], "XMAS"},
		{Coord{1, 9}, directions["NE"], "XMAS"},
		{Coord{4, 0}, directions["SE"], "XMAS"},
		{Coord{9, 3}, directions["SW"], "XMAS"},
		{Coord{9, 9}, directions["NW"], "XMAS"},
		{Coord{1, 0}, directions["SE"], "MAS"},
		{Coord{3, 0}, directions["SW"], "SAM"},
		{Coord{1, 2}, directions["NE"], "MAS"},
		{Coord{3, 2}, directions["NW"], "SAM"},
	}

	for _, test := range tests {
		actual := m.GetWord(test.pos, test.dir, len(test.word))
		if actual != test.word {
			t.Errorf("Expected %s, got %s", test.word, actual)
		}
	}
}
