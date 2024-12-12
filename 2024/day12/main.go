package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int
type Map struct {
	coords map[Coord]rune
	size   int
}

func (m *Map) Floodfill(c Coord) (int, int, int) {
	var area, perimeter, sides int

	letter := m.coords[c]
	visited := make(map[Coord]bool)
	stack := []Coord{c}

	dirs := []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	m.coords[c] = '·'
	localVisited := make(map[Coord]rune)
	localVisited[c] = '·'
	area = 0

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[curr] {
			continue
		}

		visited[curr] = true
		m.coords[curr] = '·'
		localVisited[curr] = '·'
		area++
		for _, dir := range dirs {
			next := Coord{curr[0] + dir[0], curr[1] + dir[1]}
			if visited[next] {
				continue
			}

			if m.coords[next] == letter {
				stack = append(stack, next)
				continue
			}
			perimeter++
		}
	}

	minx, miny, maxx, maxy := c[0], c[1], c[0], c[1]
	for coord := range localVisited {
		if coord[0] < minx {
			minx = coord[0]
		}
		if coord[0] > maxx {
			maxx = coord[0]
		}
		if coord[1] < miny {
			miny = coord[1]
		}
		if coord[1] > maxy {
			maxy = coord[1]
		}
	}

	for x := minx; x <= maxx+1; x++ {
		for y := miny; y <= maxy+1; y++ {
			cell := []rune{
				localVisited[Coord{x - 1, y - 1}],
				localVisited[Coord{x, y - 1}],
				localVisited[Coord{x - 1, y}],
				localVisited[Coord{x, y}],
			}
			cnt := 0
			for _, c := range cell {
				if c == '·' {
					cnt++
				}
			}
			if cnt == 1 {
				sides++
			}
			if cnt == 3 {
				sides++
			}
			if cnt == 2 && cell[0] == '·' && rune(cell[3]) == '·' {
				sides += 2
			}
			if cnt == 2 && rune(cell[1]) == '·' && rune(cell[2]) == '·' {
				sides += 2
			}
		}
	}

	return area, perimeter, sides
}

func (m Map) String() string {
	var ret string

	for y := 0; y < m.size; y++ {
		for x := 0; x < m.size; x++ {
			ret += string(m.coords[Coord{x, y}])
		}
		ret += "\n"
	}

	return ret
}

func parseInput(input string) Map {
	var ret = Map{
		coords: make(map[Coord]rune),
	}

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, c := range line {
			ret.coords[Coord{x, y}] = c
		}
		if y > ret.size {
			ret.size = y
		}
	}

	ret.size++

	return ret
}

func part1(input Map) int {
	var ret int

	for x := 0; x < input.size; x++ {
		for y := 0; y < input.size; y++ {
			if input.coords[Coord{x, y}] != '·' {
				area, perimeter, _ := input.Floodfill(Coord{x, y})
				ret += area * perimeter
			}
		}
	}

	return ret
}

func part2(input Map) int {
	var ret int

	for x := 0; x < input.size; x++ {
		for y := 0; y < input.size; y++ {
			if input.coords[Coord{x, y}] != '·' {
				area, _, sides := input.Floodfill(Coord{x, y})
				ret += area * sides
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
