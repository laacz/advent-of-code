package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

type Maze struct {
	m          [][]rune
	distances  map[Coord]int
	Start, End Coord
	Path       map[Coord]int
}

func (m Maze) String() string {
	var ret string
	for y, row := range m.m {
		for x, r := range row {
			if _, ok := m.Path[Coord{x, y}]; ok {
				ret += "·"
			} else {
				ret += string(r)
			}
		}
		ret += "\n"
	}
	return ret
}

func parseInput(input string) Maze {
	var maze Maze

	lines := strings.Split(strings.TrimSpace(input), "\n")
	maze.m = make([][]rune, len(lines))
	for y, line := range lines {
		maze.m[y] = make([]rune, len(line))
		for x, r := range line {
			maze.m[y][x] = r
			if r == 'S' {
				maze.Start = Coord{x, y}
			} else if r == 'E' {
				maze.End = Coord{x, y}
			}
		}
	}

	return maze
}

func (m *Maze) ComputeShortestPath() {
	type qItem struct {
		pos      Coord
		distance int
	}

	previous := make(map[Coord]Coord)
	visited := make(map[Coord]bool)

	queue := []qItem{{m.Start, 0}}
	m.Path = map[Coord]int{m.Start: 0}

	dirs := []Coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 {
		minIdx := 0
		for i := 1; i < len(queue); i++ {
			if queue[i].distance < queue[minIdx].distance {
				minIdx = i
			}
		}

		current := queue[minIdx]
		queue = append(queue[:minIdx], queue[minIdx+1:]...)

		if current.pos == m.End {
			break
		}

		if visited[current.pos] {
			continue
		}
		visited[current.pos] = true

		for _, dir := range dirs {
			next := Coord{current.pos.x + dir.x, current.pos.y + dir.y}

			if next.y < 0 || next.y > len(m.m)-2 || next.x < 0 || next.x > len(m.m[0])-2 ||
				m.m[next.y][next.x] == '#' || visited[next] {
				continue
			}

			newDist := m.Path[current.pos] + 1
			if currDist, exists := m.Path[next]; !exists || newDist < currDist {
				m.Path[next] = newDist
				previous[next] = current.pos
				queue = append(queue, qItem{next, newDist})
			}
		}
	}
}

func part1(maze Maze, min int) int {
	var ret int

	maze.ComputeShortestPath()

	shaveCounts := make(map[int]int)

	for pos, ps := range maze.Path {
		for _, dir := range []Coord{{0, -2}, {0, 2}, {-2, 0}, {2, 0}} {
			if nextPs, ok := maze.Path[Coord{pos.x + dir.x, pos.y + dir.y}]; ok {
				if nextPs-ps >= min {
					shaveCounts[nextPs-ps]++
				}
			}
		}
	}

	for _, count := range shaveCounts {
		ret += count
	}

	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(maze Maze, min int) int {
	var ret int

	maze.ComputeShortestPath()

	shaveCounts := make(map[int]int)

	for pos, ps := range maze.Path {
		for x := -21; x < 21; x++ {
			for y := -21; y < 21; y++ {
				if abs(x)+abs(y) > 20 {
					continue
				}

				if nextPs, ok := maze.Path[Coord{pos.x + x, pos.y + y}]; ok {
					if nextPs-ps-abs(x)-abs(y) >= min {
						shaveCounts[nextPs-ps]++
					}
				}
			}
		}

	}

	for _, count := range shaveCounts {
		ret += count
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	maze := parseInput(string(input))
	fmt.Println("part1", part1(maze, 100))
	fmt.Println("part2", part2(maze, 100))
}
