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
	dir   Direction
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

func stateKey(p Coord, d Direction) string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, d)
}

func parseInput(input string) (Maze, Coord, Coord) {
	var start, end Coord
	var maze Maze

	lines := strings.Split(strings.TrimSpace(input), "\n")
	maze = make(Maze, len(lines))

	for y, line := range lines {
		maze[y] = []rune{}
		for x, char := range line {
			if char == 'S' {
				char = '.'
				start = Coord{x, y}
			} else if char == 'E' {
				char = '.'
				end = Coord{x, y}
			}
			maze[y] = append(maze[y], char)
		}
	}

	return maze, start, end
}

func FindAllPaths(maze Maze, start, end Coord) (int, [][]Coord) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	startState := &State{
		pos:  start,
		dir:  East,
		cost: 0,
	}
	heap.Push(&pq, startState)

	visited := make(map[string]int)
	visited[stateKey(start, East)] = 0

	minCost := -1

	previous := make(map[string]map[string]State)

	moves := map[Direction][]Coord{
		North: {{0, -1}, {1, 0}, {-1, 0}},
		East:  {{1, 0}, {0, 1}, {0, -1}},
		South: {{0, 1}, {-1, 0}, {1, 0}},
		West:  {{-1, 0}, {0, -1}, {0, 1}},
	}

	newDirs := map[Direction][]Direction{
		North: {North, East, West},
		East:  {East, South, North},
		South: {South, West, East},
		West:  {West, North, South},
	}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*State)
		currentKey := stateKey(current.pos, current.dir)

		if minCost != -1 && current.cost > minCost {
			continue
		}

		if current.pos == end {
			if minCost == -1 || current.cost < minCost {
				minCost = current.cost
			}
			continue
		}

		for i, move := range moves[current.dir] {
			newPos := Coord{
				x: current.pos.x + move.x,
				y: current.pos.y + move.y,
			}
			newDir := newDirs[current.dir][i]

			if newPos.x < 0 || newPos.y < 0 ||
				newPos.y >= len(maze) || newPos.x >= len(maze[0]) ||
				maze[newPos.y][newPos.x] == '#' {
				continue
			}

			newCost := current.cost + 1
			if newDir != current.dir {
				newCost += 1000
			}

			newStateKey := stateKey(newPos, newDir)

			if cost, exists := visited[newStateKey]; !exists || newCost <= cost {
				if !exists || newCost < cost {
					visited[newStateKey] = newCost
					previous[newStateKey] = make(map[string]State)
				}

				if newCost == visited[newStateKey] {
					if previous[newStateKey] == nil {
						previous[newStateKey] = make(map[string]State)
					}
					previous[newStateKey][currentKey] = *current

					heap.Push(&pq, &State{
						pos:  newPos,
						dir:  newDir,
						cost: newCost,
					})
				}
			}
		}
	}

	if minCost == -1 {
		return -1, nil
	}

	var allPaths [][]Coord
	seen := make(map[string]bool)

	var buildPath func(current State, path []Coord)
	buildPath = func(current State, path []Coord) {
		if current.pos == start {
			newPath := make([]Coord, len(path))
			copy(newPath, path)
			allPaths = append(allPaths, newPath)
			return
		}

		currentKey := stateKey(current.pos, current.dir)
		if seen[currentKey] {
			return
		}
		seen[currentKey] = true

		for _, prevState := range previous[currentKey] {
			newPath := make([]Coord, 0, len(path)+1)
			newPath = append(newPath, current.pos)
			newPath = append(newPath, path...)
			buildPath(prevState, newPath)
		}

		seen[currentKey] = false
	}

	for dir := Direction(0); dir < 4; dir++ {
		endStateKey := stateKey(end, dir)
		if cost, exists := visited[endStateKey]; exists && cost == minCost {
			buildPath(State{pos: end, dir: dir, cost: minCost}, nil)
		}
	}

	return minCost, allPaths
}

func part1(maze Maze, start, exit Coord) int {
	ret, _ := FindAllPaths(maze, start, exit)

	return ret
}

func part2(maze Maze, start, exit Coord) int {
	var ret int

	_, paths := FindAllPaths(maze, start, exit)
	seen := make(map[Coord]bool)
	for _, p := range paths {
		for _, pt := range p {
			if !seen[pt] {
				seen[pt] = true
				ret++
			}
		}
	}

	return ret + 1
}

func main() {
	input, _ := os.ReadFile("input.txt")

	maze, start, end := parseInput(string(input))
	fmt.Println("part1", part1(maze, start, end))
	fmt.Println("part2", part2(maze, start, end))
}
