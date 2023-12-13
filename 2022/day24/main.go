package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Direction Point

type Blizzard struct {
	pos Point
	dir Direction
}

func (b Blizzard) String() string {
	switch b.dir {
	case Up:
		return "^"
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	default:
		return "?"
	}
}

var Up = Direction{0, -1}
var Down = Direction{0, 1}
var Left = Direction{-1, 0}
var Right = Direction{1, 0}

var Directions = []Direction{Up, Down, Left, Right}

type Map struct {
	blizzards     []*Blizzard
	blizzardsMap  map[Point]bool
	width, height int
}

// NewMap creates a new map from the input
func NewMap(input string) Map {
	lines := strings.Split(input, "\n")
	m := Map{
		width:  len(lines[0]),
		height: len(lines),
	}

	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '>':
				m.blizzards = append(m.blizzards, &Blizzard{Point{x, y}, Right})
			case '<':
				m.blizzards = append(m.blizzards, &Blizzard{Point{x, y}, Left})
			case '^':
				m.blizzards = append(m.blizzards, &Blizzard{Point{x, y}, Up})
			case 'v':
				m.blizzards = append(m.blizzards, &Blizzard{Point{x, y}, Down})
			default:
				// do nothing
			}
		}
	}

	return m
}

// Move moves the blizzard one step in its direction and/or wraps it around the map
func (m *Map) Move(b *Blizzard) {
	b.pos.x += b.dir.x
	b.pos.y += b.dir.y

	if b.pos.x < 1 || b.pos.x > m.width-2 || b.pos.y < 1 || b.pos.y > m.height-2 {
		switch b.dir {
		case Up:
			b.pos = Point{b.pos.x, m.height - 2}
		case Down:
			b.pos = Point{b.pos.x, 1}
		case Left:
			b.pos = Point{m.width - 2, b.pos.y}
		case Right:
			b.pos = Point{1, b.pos.y}
		}
	}
}

// UpdateBlizzardsPos updates the blizzards map for faster access
func (m *Map) UpdateBlizzardsPos() {
	m.blizzardsMap = make(map[Point]bool)
	for _, b := range m.blizzards {
		m.blizzardsMap[b.pos] = true
	}
}

// Solve fills the map from start to the end and returns the number of steps (minutes)
func (m *Map) Solve(from, to Point) int {
	time := 0
	filled := make(map[Point]bool)
	filled[from] = true

	// ok, flood fill with a catch (a blizzard clears already filled points)
	for !filled[to] {
		time++
		for _, b := range m.blizzards {
			m.Move(b)
		}
		m.UpdateBlizzardsPos()

		filled2 := make(map[Point]bool)
		// re-iterate over all filled points
		for p := range filled {
			// if empty, fill
			if _, ok := m.blizzardsMap[p]; !ok {
				filled2[p] = true
			}
			// fill all free at all directions
			for _, dir := range Directions {
				p2 := Point{p.x + dir.x, p.y + dir.y}

				_, ok2 := m.blizzardsMap[p2]
				if (!ok2 && p2.x > 0 && p2.x < m.width-1 && p2.y > 0 && p2.y < m.height-1) ||
					p2 == to {
					filled2[p2] = true
				}
			}
		}
		filled = filled2
	}

	return time
}

func partOne(input string) int {
	m := NewMap(input)
	return m.Solve(Point{1, 0}, Point{m.width - 2, m.height - 1})
}

func partTwo(input string) int {
	m := NewMap(input)
	return m.Solve(Point{1, 0}, Point{m.width - 2, m.height - 1}) +
		m.Solve(Point{m.width - 2, m.height - 1}, Point{1, 0}) +
		m.Solve(Point{1, 0}, Point{m.width - 2, m.height - 1})
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
