package main

import (
	"fmt"
	"math"

	"github.com/laacz/aoc-2016/util"
)

type Point struct {
	x, y int
}

type Grid struct {
	cells  [][]byte
	pois   map[int]Point
	width  int
	height int
}

func parse(input string) Grid {
	lines := util.GetLines(input)
	grid := Grid{
		cells: make([][]byte, len(lines)),
		pois:  make(map[int]Point),
	}
	grid.height = len(lines)

	for y, line := range lines {
		grid.cells[y] = []byte(line)
		if y == 0 {
			grid.width = len(line)
		}
		for x, ch := range line {
			if ch >= '0' && ch <= '9' {
				grid.pois[int(ch-'0')] = Point{x, y}
			}
		}
	}
	return grid
}

func (g *Grid) bfs(from, to Point) int {
	type State struct {
		pos   Point
		steps int
	}

	visited := make(map[Point]bool)
	queue := []State{{from, 0}}
	visited[from] = true

	dirs := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.pos == to {
			return curr.steps
		}

		for _, d := range dirs {
			next := Point{curr.pos.x + d.x, curr.pos.y + d.y}
			if next.x < 0 || next.x >= g.width || next.y < 0 || next.y >= g.height {
				continue
			}
			if g.cells[next.y][next.x] == '#' {
				continue
			}
			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, State{next, curr.steps + 1})
		}
	}
	return -1
}

func (g *Grid) distances() map[int]map[int]int {
	distances := make(map[int]map[int]int)
	for i := range g.pois {
		distances[i] = make(map[int]int)
	}

	for i, pi := range g.pois {
		for j, pj := range g.pois {
			if i < j {
				dist := g.bfs(pi, pj)
				distances[i][j] = dist
				distances[j][i] = dist
			}
		}
	}
	return distances
}

func permute(pois []int) [][]int {
	if len(pois) == 0 {
		return [][]int{{}}
	}

	var result [][]int
	for i, poi := range pois {
		rest := make([]int, 0, len(pois)-1)
		rest = append(rest, pois[:i]...)
		rest = append(rest, pois[i+1:]...)

		for _, perm := range permute(rest) {
			result = append(result, append([]int{poi}, perm...))
		}
	}
	return result
}

func partOne(input string) int {
	grid := parse(input)
	distances := grid.distances()

	var otherPOIs []int
	for poi := range grid.pois {
		if poi != 0 {
			otherPOIs = append(otherPOIs, poi)
		}
	}

	minSteps := math.MaxInt32
	for _, perm := range permute(otherPOIs) {
		steps := distances[0][perm[0]]
		for i := 0; i < len(perm)-1; i++ {
			steps += distances[perm[i]][perm[i+1]]
		}
		if steps < minSteps {
			minSteps = steps
		}
	}

	return minSteps
}

func partTwo(input string) int {
	grid := parse(input)
	distances := grid.distances()

	var otherPOIs []int
	for poi := range grid.pois {
		if poi != 0 {
			otherPOIs = append(otherPOIs, poi)
		}
	}

	minSteps := math.MaxInt32
	for _, perm := range permute(otherPOIs) {
		steps := distances[0][perm[0]]
		for i := 0; i < len(perm)-1; i++ {
			steps += distances[perm[i]][perm[i+1]]
		}
		steps += distances[perm[len(perm)-1]][0]
		if steps < minSteps {
			minSteps = steps
		}
	}

	return minSteps
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
