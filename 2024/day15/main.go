package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Map struct {
	cells     [][]byte
	width     int
	height    int
	movements string
	moveNum   int
	pos       []int
	isDobled  bool
	lastMove  byte
}

func (m *Map) String() string {
	var ret string
	for y := range m.height + 1 {
		for x := range m.width + 1 {
			c := m.cells[y][x]
			if m.pos[0] == x && m.pos[1] == y {
				ret += "\033[103;30m"
				if m.lastMove == 0 {
					c = '@'
				} else {
					c = m.lastMove
				}
			} else if c == '#' {
				ret += "\033[31m"
			} else if c == 'O' || c == '[' || c == ']' {
				ret += "\033[32m"
			}

			ret += string(c)
			ret += "\033[0m"
		}
		ret += "\n"
	}

	return strings.TrimSpace(ret)
}

func parseInput(input string) *Map {
	var m Map

	parts := strings.Split(input, "\n\n")
	m.cells = make([][]byte, len(parts[0]))
	for y, line := range strings.Split(parts[0], "\n") {
		for x, c := range line {
			if c == '@' {
				m.pos = []int{x, y}
				line = strings.Replace(line, "@", " ", 1)
			}
		}
		line = strings.Replace(line, ".", " ", -1)
		m.cells[y] = []byte(line)
		m.width = y
		m.height = y
	}

	for _, c := range parts[1] {
		switch c {
		case '^', '>', 'v', '<':
			m.movements += string(c)
		}
	}

	return &m
}

func (m *Map) Resize() {
	var newCells [][]byte

	for _, row := range m.cells {
		var newRow []byte
		for _, cell := range row {
			switch cell {
			case '#':
				newRow = append(newRow, "##"...)
			case 'O':
				newRow = append(newRow, "[]"...)
			case ' ':
				newRow = append(newRow, "  "...)
			case '@':
				newRow = append(newRow, "@."...)
			default:
				panic(fmt.Sprintf("Unknown cell [%c]", cell))
			}
		}
		newCells = append(newCells, newRow)
	}

	m.cells = newCells
	m.width = m.width*2 + 1
	m.pos[0] = m.pos[0] * 2
}

func (m *Map) MoveX(move int) {
	x, y := m.pos[0], m.pos[1]
	x += move
	for m.cells[y][x] != '#' && m.cells[y][x] != ' ' {
		x += move
	}

	if m.cells[y][x] == ' ' {
		for x != m.pos[0] || y != m.pos[1] {
			m.cells[y][x] = m.cells[y][x-move]
			x -= move
		}
		m.cells[m.pos[1]][m.pos[0]] = ' '
		m.pos[0] = x + move
	}
}

func (m *Map) MoveY(move int) {
	y := m.pos[1]

	boxes := [][]int{}

	if m.cells[m.pos[1]+move][m.pos[0]] == '#' {
		return
	}

	coords := []int{m.pos[0]}
	for {
		y += move
		newCoords := []int{}
		for _, c := range coords {
			if m.cells[y][c] == '[' {
				newCoords = append(newCoords, c, c+1)
				boxes = append(boxes, []int{c, y})
			}
			if m.cells[y][c] == ']' {
				boxes = append(boxes, []int{c - 1, y})
				newCoords = append(newCoords, c-1, c)
			}
			if m.cells[y][c] == 'O' {
				boxes = append(boxes, []int{c, y})
				newCoords = append(newCoords, c)
			}
		}
		if len(newCoords) == 0 {
			break
		}
		coords = newCoords
	}

	// I couldn't be bothered to improve upon previous code, so let's
	// deal with consequences, not causes.
	seen := map[[2]int]struct{}{}
	newBoxes := [][]int{}
	for _, box := range boxes {
		if _, ok := seen[[2]int{box[0], box[1]}]; !ok {
			seen[[2]int{box[0], box[1]}] = struct{}{}
			newBoxes = append(newBoxes, box)
		}
	}

	boxes = newBoxes

	slices.Reverse(boxes)
	for i := range boxes {
		if m.cells[boxes[i][1]+move][boxes[i][0]] == '#' {
			return
		}
		if m.width != m.height {
			if m.cells[boxes[i][1]+move][boxes[i][0]+1] == '#' {
				return
			}
		}
	}

	m.cells[m.pos[1]][m.pos[0]] = ' '

	for i := range boxes {
		symbol := m.cells[boxes[i][1]][boxes[i][0]]
		m.cells[boxes[i][1]+move][boxes[i][0]] = symbol
		m.cells[boxes[i][1]][boxes[i][0]] = ' '

		if m.width != m.height {
			m.cells[boxes[i][1]+move][boxes[i][0]+1] = ']'
			m.cells[boxes[i][1]][boxes[i][0]+1] = ' '
		}
	}

	m.pos[1] += move
}

func (m *Map) Move() {
	move := m.movements[m.moveNum]
	m.lastMove = m.movements[m.moveNum]
	m.moveNum++

	switch move {
	case '^':
		m.MoveY(-1)
	case '>':
		m.MoveX(1)
	case 'v':
		m.MoveY(1)
	case '<':
		m.MoveX(-1)
	}
}

func part1(m *Map) int {
	var ret int

	for range len(m.movements) {
		m.Move()
	}
	fmt.Println(m)

	for y := range m.height + 1 {
		for x := range m.width + 1 {
			if m.cells[y][x] == 'O' {
				ret += y*100 + x
			}
		}
	}

	return ret
}
func part2(m *Map) int {
	var ret int

	m.Resize()

	for range len(m.movements) {
		m.Move()
	}
	fmt.Println(m)

	for y := range m.height + 1 {
		for x := range m.width + 1 {
			if m.cells[y][x] == '[' {
				ret += y*100 + x
			}
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
