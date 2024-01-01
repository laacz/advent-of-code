package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input     string
	expected1 int
	expected2 int
}{
	{
		input: `
		#.#####################
		#.......#########...###
		#######.#########.#.###
		###.....#.>.>.###.#.###
		###v#####.#v#.###.#.###
		###.>...#.#.#.....#...#
		###v###.#.#.#########.#
		###...#.#.#.......#...#
		#####.#.#.#######.#.###
		#.....#.#.#.......#...#
		#.#####.#.#.#########v#
		#.#...#...#...###...>.#
		#.#.#v#######v###.###v#
		#...#.>.#...>.>.#.###.#
		#####v#.#.###v#.#.###.#
		#.....#...#...#.#.#...#
		#.#########.###.#.#.###
		#...###...#...#...#.###
		###.###.#.###v#####v###
		#...#...#.#.>.>.#.>.###
		#.###.###.#.###.#.#v###
		#.....###...###...#...#
		#####################.#`,
		expected1: 94,
		expected2: 154,
	},
}

func contains(s []Coord, c Coord) bool {
	for _, e := range s {
		if e == c {
			return true
		}
	}
	return false
}

func TestParse(t *testing.T) {
	m, graph := parse(tests[0].input)

	l := len(m)

	celltests := []struct {
		c        Coord
		expected []Coord
	}{
		{
			c:        Coord{1, 0},
			expected: []Coord{{1, 1}},
		},
		{
			c:        Coord{l - 2, l - 1},
			expected: []Coord{{l - 2, l - 2}},
		},
		{
			c:        Coord{1, 1},
			expected: []Coord{{1, 0}, {2, 1}},
		},
	}

	for _, tt := range celltests {
		if len(graph[tt.c]) != len(tt.expected) {
			t.Errorf("Expected %d neighbours for {%d, %d}, got %d", len(tt.expected), tt.c[0], tt.c[1], len(graph[tt.c]))
			fmt.Println(graph[tt.c])
		}
		for c := range graph[tt.c] {
			if !contains(tt.expected, c) {
				t.Errorf("Expected %v, got %v", tt.expected, graph[tt.c])
			}
		}
	}

	if len(graph) != 213 {
		t.Errorf("Expected 213 nodes, got %d", len(graph))
	}

	if graph[Coord{0, 0}] != nil {
		t.Errorf("Expected nil, got %v", graph[Coord{0, 0}])
	}
}

func TestContract(t *testing.T) {
	m, graph := parse(tests[0].input)
	graph.Contract()

	l := len(m)

	if len(graph[Coord{1, 0}]) != 1 {
		t.Errorf("Expected 1 neighbour for {1, 0}, got %d", len(graph[Coord{1, 0}]))
	}

	if len(graph[Coord{l - 2, l - 1}]) != 1 {
		t.Errorf("Expected 1 neighbour for {%d, %d}, got %d", l-2, l-1, len(graph[Coord{1, 0}]))
	}

	if graph[Coord{0, 0}] != nil {
		t.Errorf("Expected nil, got %v", graph[Coord{0, 0}])
	}

	if len(graph) != 9 {
		t.Errorf("Expected 9 nodes, got %d", len(graph))
	}

}

func TestContractWeights(t *testing.T) {
	_, graph := parse(tests[0].input)
	graph.Contract()

	for _, tt := range []struct {
		c1       Coord
		c2       Coord
		expected int
	}{
		{c1: Coord{1, 0}, c2: Coord{3, 5}, expected: 15},
		{c1: Coord{5, 13}, c2: Coord{3, 5}, expected: 22},
		{c1: Coord{5, 13}, c2: Coord{13, 19}, expected: 38},
		{c1: Coord{5, 13}, c2: Coord{13, 13}, expected: 12},
		{c1: Coord{11, 3}, c2: Coord{3, 5}, expected: 22},
		{c1: Coord{21, 22}, c2: Coord{19, 19}, expected: 5},
	} {
		actual := graph[tt.c1][tt.c2]
		if actual != tt.expected {
			t.Errorf("Expected %d from %v to %v, got %d", tt.expected, tt.c1, tt.c2, actual)
		}
	}
}

func TestPartOne(t *testing.T) {
	for _, tt := range tests {
		actual := partOne(tt.input)
		if actual != tt.expected1 {
			t.Errorf("Expected %d, got %d", tt.expected1, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, tt := range tests {
		actual := partTwo(tt.input)
		if actual != tt.expected2 {
			t.Errorf("Expected %d, got %d", tt.expected2, actual)
		}
	}
}
