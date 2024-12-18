package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Maze [][]rune

func (m Maze) String() string {
	var ret string
	for _, row := range m {
		ret += string(row) + "\n"
	}
	return ret
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Coord struct {
	x, y int
}

type State struct {
	pos   Coord
	cost  int
	index int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	state := x.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	old[n-1] = nil
	state.index = -1
	*pq = old[0 : n-1]
	return state
}

func stateKey(p Coord) string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y)
}

func parseInput(input string, size int, limit int) (Maze, []Coord) {
	var maze Maze
	var bytes []Coord

	lines := strings.Split(strings.TrimSpace(input), "\n")
	maze = make(Maze, size)
	for y := range maze {
		maze[y] = []rune(strings.Repeat(".", size))
	}

	for i, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		maze[y][x] = '#'
		bytes = append(bytes, Coord{x, y})
		if limit != 0 && i > limit {
			break
		}
	}

	return maze, bytes
}

func FindShortestPath(maze Maze) int {
	start := Coord{0, 0}
	end := Coord{len(maze[0]) - 1, len(maze) - 1}
	directions := []Coord{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{pos: start, cost: 0})

	costs := make(map[Coord]int)
	costs[start] = 0

	paths := make(map[Coord][]Coord)
	paths[start] = []Coord{start}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)
		if current.pos == end {
			return current.cost
		}

		for _, d := range directions {
			next := Coord{current.pos.x + d.x, current.pos.y + d.y}
			if next.x < 0 || next.x >= len(maze[0]) || next.y < 0 || next.y >= len(maze) || maze[next.y][next.x] == '#' {
				continue
			}

			newCost := current.cost + 1
			if c, ok := costs[next]; !ok || newCost < c {
				costs[next] = newCost
				heap.Push(pq, &State{pos: next, cost: newCost})
				paths[next] = append(paths[current.pos], next)
			}
		}
	}

	return -1
}

func part1(maze Maze) int {
	ret := FindShortestPath(maze)

	return ret
}

func part2(input string) string {
	var i = 1023
	var ret string

	for {
		i++
		maze, bytes := parseInput(string(input), 71, i)
		pathLen := FindShortestPath(maze)
		if pathLen == -1 {
			ret = fmt.Sprintf("%d,%d", bytes[len(bytes)-1].x, bytes[len(bytes)-1].y)
			break
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	maze, _ := parseInput(string(input), 71, 1023)
	fmt.Println("part1", part1(maze))
	fmt.Println("part2", part2(string(input)))
}
