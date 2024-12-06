package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type XY [2]int

type Map struct {
	cells    map[XY]byte
	pos      XY
	dir      XY
	size     XY
	wallsHit map[XY][]XY
}

var dirChars = map[XY]byte{
	{1, 0}:  '>',
	{-1, 0}: '<',
	{0, 1}:  'v',
	{0, -1}: '^',
}

func (m Map) String() string {
	var ret strings.Builder

	for y := 0; y <= m.size[1]; y++ {
		for x := 0; x <= m.size[0]; x++ {
			if x == m.pos[0] && y == m.pos[1] {
				ret.WriteByte(dirChars[m.dir])
			} else {
				ret.WriteByte(m.cells[XY{x, y}])
			}
		}
		ret.WriteByte('\n')
	}

	return ret.String()
}

func (m Map) Count() int {
	var ret int
	for _, c := range m.cells {
		if c == 'X' {
			ret++
		}
	}

	return ret
}

// Walk returns true if the robot is stuck in a loop
func (m *Map) Walk() bool {
	for {
		if m.pos[0] < 0 || m.pos[1] < 0 || m.pos[0] > m.size[0] || m.pos[1] > m.size[1] {
			break
		}

		m.cells[m.pos] = 'X'

		for m.cells[XY{m.pos[0] + m.dir[0], m.pos[1] + m.dir[1]}] == '#' {
			if slices.Contains(m.wallsHit[m.pos], m.dir) {
				return true
			}
			m.wallsHit[m.pos] = append(m.wallsHit[m.pos], m.dir)
			m.dir[0], m.dir[1] = -m.dir[1], m.dir[0]
		}

		m.pos[0] += m.dir[0]
		m.pos[1] += m.dir[1]
	}

	return false
}

func parseInput(input string) Map {
	ret := Map{
		cells:    make(map[XY]byte),
		dir:      XY{0, -1},
		wallsHit: make(map[XY][]XY),
	}
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, c := range line {
			if c == '^' {
				ret.pos = XY{x, y}
				c = '.'
			}
			ret.cells[XY{x, y}] = byte(c)
			ret.size[0] = x
		}
		ret.size[1] = y
	}

	return ret
}

func part1(input string) int {
	m := parseInput(input)
	m.Walk()

	return m.Count()
}

func part2(input string) int {
	var ret int
	m := parseInput(input)
	m.Walk()
	positions := make(map[XY]bool)
	for xy, c := range m.cells {
		if c == 'X' {
			positions[xy] = true
		}
	}

	for xy := range positions {
		m := parseInput(input)
		m.cells[xy] = '#'
		if m.Walk() {
			ret++
		}
	}

	return ret
}

func main() {

	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
