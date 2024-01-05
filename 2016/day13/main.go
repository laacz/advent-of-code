package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

type Coords struct {
	x, y int
}

type Map struct {
	magic      int
	goal       Coords
	visited    map[Coords]bool
	maxVisited Coords
}

func (m Map) String() (ret string) {
	ret = "  "
	for i := 0; i < m.maxVisited.x+1; i++ {
		ret += fmt.Sprintf("%d", i%10)
	}
	ret += "\n"
	for y := 0; y <= m.maxVisited.y; y++ {
		ret += fmt.Sprintf("%d ", y%10)
		for x := 0; x <= m.maxVisited.x; x++ {
			if m.goal.x == x && m.goal.y == y {
				// ansi escape sequence bright white background and black text
				ret += "\033[47;30m"
			}
			if m.visited[Coords{x, y}] {
				ret += "O"
			} else if m.IsWall(Coords{x, y}) {
				ret += "#"
			} else {
				ret += "."
			}

			// ansi reset
			ret += "\033[0m"
		}
		ret += "\n"
	}

	return
}

func (m Map) IsWall(pos Coords) bool {
	x := pos.x
	y := pos.y
	if x < 0 || y < 0 {
		return true
	}
	n := x*x + 3*x + 2*x*y + y + y*y + m.magic
	// count the number of bits that are 1
	bits := 0
	for n > 0 {
		bits += n & 1
		n >>= 1
	}
	return bits%2 == 1
}

func parse(input string) (ret Map) {
	lines := util.GetLines(input)
	ret = Map{}
	ret.visited = make(map[Coords]bool)
	fmt.Sscanf(lines[0], "%d %d %d", &ret.magic, &ret.goal.x, &ret.goal.y)
	return
}

func (m *Map) bfs(start Coords) (ret [][]int) {
	queue := []Coords{start}
	visited := make(map[Coords]bool)
	parent := make(map[Coords]Coords)
	visited[Coords{start.x, start.y}] = true

	dirs := []Coords{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Directions: Right, Down, Left, Up

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == m.goal {
			return constructPath(current, parent)
		}

		for _, dir := range dirs {
			next := Coords{current.x + dir.x, current.y + dir.y}
			if !m.IsWall(next) && !visited[next] {
				parent[next] = current
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return nil // No path found
}

func (m *Map) bfs2(start Coords, visited map[Coords]bool) {
	queue := []Coords{start}
	parent := make(map[Coords]Coords)
	visited[Coords{start.x, start.y}] = true

	dirs := []Coords{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Directions: Right, Down, Left, Up

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(constructPath(current, parent)) > 49 {
			return
		}

		for _, dir := range dirs {
			next := Coords{current.x + dir.x, current.y + dir.y}
			if !m.IsWall(next) && !visited[next] {
				parent[next] = current
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
}

func constructPath(goal Coords, parent map[Coords]Coords) [][]int {
	var path [][]int
	for at, exists := goal, true; exists; at, exists = parent[at] {
		path = append([][]int{{at.x, at.y}}, path...)
	}
	return path
}

func partOne(input string) (ret int) {
	m := parse(input)

	p := m.bfs(Coords{1, 1})
	// -1 since the start doesn't count
	ret = len(p) - 1

	return
}

func partTwo(input string) (ret int) {
	m := parse(input)

	v := make(map[Coords]bool)
	m.bfs2(Coords{1, 1}, v)
	// +1 since the start counts
	ret = len(v) + 1

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
