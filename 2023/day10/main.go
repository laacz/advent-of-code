package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coords [2]int
type Map struct {
	field   map[Coords]Pipe
	start   Coords
	visited map[Coords]bool
	w, h    int
}

// String pretty-prints a Map
func (m Map) String() string {
	ret := ""

	for y := 0; y < m.h+1; y++ {
		for x := 0; x < m.w+1; x++ {
			mm := m.field[Coords{y, x}]
			c := mm.Type
			color := "\033[38;2;32;32;32m"
			if m.visited[Coords{y, x}] {
				color = "\033[38;2;255;255;0m"
			} else {
				if mm.Inside {
					color = "\033[38;2;0;255;0m"
					c = 'I'
				} else {
					color = "\033[38;2;0;0;255m"
				}
			}
			switch c {
			case '7':
				c = '┐'
			case 'F':
				c = '┌'
			case 'L':
				c = '└'
			case 'J':
				c = '┘'
			case '-':
				c = '─'
			case '|':
				c = '│'
			case '.':
				c = '⊙'
			}
			if y == m.start[0] && x == m.start[1] {
				// add light gray background to start
				color = "\033[48;2;0;0;0m"
			}

			ret = ret + color + string(c) + "\033[0m"
		}
		ret = ret + "\n"
	}

	return ret
}

type Pipe struct {
	Type        rune
	Distance    int
	Connections []Coords
	Inside      bool
}

// parse (you won't believe it!) returns parsed input
func parse(input string) Map {
	ret := Map{
		field:   make(map[Coords]Pipe),
		visited: make(map[Coords]bool),
	}

	for y, line := range strings.Split(input, "\n") {

		if len(line) == 0 {
			continue
		}
		line = strings.TrimSpace(line)

		// Now let's build a list of pipes and their connections
		for x, c := range line {
			pipe := Pipe{Type: c}

			switch c {
			case '|':
				pipe.Connections = []Coords{{y - 1, x}, {y + 1, x}}
			case '-':
				pipe.Connections = []Coords{{y, x - 1}, {y, x + 1}}
			case 'L':
				pipe.Connections = []Coords{{y, x + 1}, {y - 1, x}}
			case 'J':
				pipe.Connections = []Coords{{y, x - 1}, {y - 1, x}}
			case '7':
				pipe.Connections = []Coords{{y, x - 1}, {y + 1, x}}
			case 'F':
				pipe.Connections = []Coords{{y, x + 1}, {y + 1, x}}
			case 'S':
				ret.start = Coords{y, x}
			}

			ret.field[Coords{y, x}] = pipe

			if ret.h < y {
				ret.h = y
			}

			if ret.w < x {
				ret.w = x
			}
		}
	}

	s := ret.field[ret.start]
	sy := ret.start[0]
	sx := ret.start[1]

	// Starting pipe connections are inferred, not given
	if strings.Contains("|F7", string(ret.field[Coords{sy - 1, sx}].Type)) {
		s.Connections = append(s.Connections, Coords{sy - 1, sx})
	}
	if strings.Contains("-FL", string(ret.field[Coords{sy, sx - 1}].Type)) {
		s.Connections = append(s.Connections, Coords{sy, sx - 1})
	}
	if strings.Contains("|JL", string(ret.field[Coords{sy + 1, sx}].Type)) {
		s.Connections = append(s.Connections, Coords{sy + 1, sx})
	}
	if strings.Contains("-J7", string(ret.field[Coords{sy, sx + 1}].Type)) {
		s.Connections = append(s.Connections, Coords{sy, sx + 1})
	}

	no := Coords{sy - 1, sx}
	so := Coords{sy + 1, sx}
	we := Coords{sy, sx - 1}
	ea := Coords{sy, sx + 1}

	// Now we can infer the type of the starting pipe
	if slices.Contains(s.Connections, no) && slices.Contains(s.Connections, so) {
		s.Type = '|'
	} else if slices.Contains(s.Connections, we) && slices.Contains(s.Connections, ea) {
		s.Type = '-'
	} else if slices.Contains(s.Connections, we) && slices.Contains(s.Connections, so) {
		s.Type = '7'
	} else if slices.Contains(s.Connections, we) && slices.Contains(s.Connections, no) {
		s.Type = 'J'
	} else if slices.Contains(s.Connections, ea) && slices.Contains(s.Connections, no) {
		s.Type = 'L'
	} else if slices.Contains(s.Connections, ea) && slices.Contains(s.Connections, so) {
		s.Type = 'F'
	} else {
		panic("Unknown pipe type")
	}

	ret.field[Coords{sy, sx}] = s
	return ret
}

func (m Map) Start() (int, int) {
	for k, v := range m.field {
		if v.Type == 'S' {
			return k[0], k[1]
		}
	}
	return 0, 0
}

func (m Map) Walk() {
	// Let's try a simple DFS, going both ways simlutaniously
	queue := []Coords{m.start}
	i := 0

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		pipe := m.field[c]

		i += 1

		for _, conn := range pipe.Connections {
			if m.visited[conn] {
				continue
			}
			m.visited[conn] = true
			c := m.field[conn]
			c.Distance = pipe.Distance + 1
			m.field[conn] = c
			queue = append(queue, conn)
		}
	}
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 0

	m := parse(input)
	m.Walk()

	// Find the futhest distance and we're done
	for _, v := range m.field {
		if sum < v.Distance {
			sum = v.Distance
		}
	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	sum := 0

	m := parse(input)
	m.Walk()

	for c := range m.field {
		if m.visited[c] {
			continue
		}
		n := 0

		// Now for that little sweet and magic winding number approach for determining if a point is inside
		// a polygon or not (as the path forms a closed polygon).
		for x := c[1]; x <= m.w; x++ {
			if strings.Contains("|LJ", string(m.field[Coords{c[0], x}].Type)) && m.visited[Coords{c[0], x}] {
				n += 1
			}
		}
		if n%2 == 1 {
			f := m.field[c]
			f.Inside = true
			m.field[c] = f
			sum += 1
		}
	}

	fmt.Println(m)

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
