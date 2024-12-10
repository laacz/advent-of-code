package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int

type Map map[Coord]int

func (m Map) Size() int {
	size := 0
	for c := range m {
		if c[0] > size {
			size = c[0]
		}
	}
	return size
}

func (m Map) ReachableNines(start Coord) int {
	visited := make(map[Coord]bool)
	var ret int

	var directions = []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var size = m.Size() + 1
	var queue = []Coord{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		val := m[current]

		for _, d := range directions {
			next := Coord{current[0] + d[0], current[1] + d[1]}

			if next[0] < 0 || next[1] < 0 || next[0] >= size || next[1] >= size {
				continue
			}

			nextVal := m[next]
			if nextVal == val+1 {
				queue = append(queue, next)
				if nextVal == 9 && !visited[next] {
					ret++
					visited[next] = true
				}
			}
		}
	}

	return ret
}

func (m Map) AllPaths(start Coord) int {
	visited := make(map[Coord]bool)
	var ret int

	var directions = []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var size = m.Size() + 1
	var queue = []Coord{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current] = true
		val := m[current]

		for _, d := range directions {
			next := Coord{current[0] + d[0], current[1] + d[1]}

			if next[0] < 0 || next[1] < 0 || next[0] >= size || next[1] >= size {
				continue
			}

			nextVal := m[next]
			if nextVal == val+1 {
				queue = append(queue, next)
				if nextVal == 9 {
					ret++
				}
			}
		}
	}

	return ret
}

func part1(input Map) int {
	var ret int

	var size = input.Size()
	for y := range size + 1 {
		for x := range size + 1 {
			if input[Coord{x, y}] == 0 {
				ret += input.ReachableNines(Coord{x, y})
			}
		}
	}

	return ret
}

func part2(input Map) int {
	var ret int

	var size = input.Size()
	for y := range size + 1 {
		for x := range size + 1 {
			if input[Coord{x, y}] == 0 {
				ret += input.AllPaths(Coord{x, y})
			}
		}
	}

	return ret
}

func parseInput(input string) Map {
	ret := make(Map)
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, c := range line {
			ret[Coord{x, y}] = int(c - '0')
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
