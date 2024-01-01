package main

import (
	"fmt"
	"os"
	"strings"
)

var path []Coord
var m Map
var graph Graph

type Coord [2]int
type Map [][]rune

func (m Map) End() (ret Coord) {
	return Coord{len(m[0]) - 2, len(m) - 1}
}

func (c Coord) InSlice(s []Coord) bool {
	for _, e := range s {
		if e == c {
			return true
		}
	}
	return false
}

func (m Map) String() string {
	var ret string
	for y, row := range m {
		for x, c := range row {
			coo := Coord{x, y}
			if graph[coo] != nil {
				ret += "\033[1;30;42m"
			}
			if coo.InSlice(path) && c == '.' {
				ret += "\033[1;34m"
				ret += "O"
			} else {
				if coo.InSlice(path) && c != '.' {
					ret += "\033[1;30;44m"
				}
				switch c {
				case '>':
					ret += "→"
				case '<':
					ret += "←"
				case '^':
					ret += "↑"
				case 'v':
					ret += "↓"
				default:
					ret += string(c)
				}
			}
			ret += "\033[0m"
		}
		ret += "\n"
	}
	return ret
}

type Graph map[Coord]map[Coord]int

// Contract removes nodes with only two edges and merging remainig edges
func (g *Graph) Contract() {
	for k, v := range *g {
		if len(v) == 2 {
			left := Coord{}
			right := Coord{}

			for c := range v {
				if left[0] == 0 {
					left = c
				} else {
					right = c
				}
			}

			// remove k from left and right
			delete((*g)[left], k)
			delete((*g)[right], k)

			// add left to right
			(*g)[left][right] = v[left] + v[right]
			(*g)[right][left] = v[left] + v[right]

			// remove k
			delete(*g, k)
		}
	}
}

// DFS implements classic DFS
func (g Graph) DFS(start, end Coord, visited map[Coord]bool, slipperySlope bool) (ret []int) {

	if start == end {
		return []int{0}
	}

	visited[start] = true

	x, y := start[0], start[1]

	for c := range g[start] {
		nx, ny := c[0], c[1]

		if slipperySlope {
			if (m[y][x] == '>' && (nx != x+1 || ny != y)) ||
				(m[y][x] == 'v' && (nx != x || ny != y+1)) ||
				(m[y][x] == '<' && (nx != x-1 || ny != y)) ||
				(m[y][x] == '^' && (nx != x || ny != y-1)) {
				continue
			}
		}

		if visited[c] {
			continue
		}

		paths := g.DFS(c, m.End(), visited, slipperySlope)

		for _, l := range paths {
			ret = append(ret, l+g[start][c])
		}
	}

	delete(visited, start)

	return ret
}

func parse(input string) (m Map, ret Graph) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var row []rune
		for _, c := range line {
			var num int
			fmt.Sscanf(string(c), "%d", &num)
			row = append(row, c)
		}
		m = append(m, row)
	}

	// now build graph
	ret = make(Graph)
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == '#' {
				continue
			}

			// add all possible directions
			for _, dir := range []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				ny, nx := y+dir[0], x+dir[1]

				if ny < 0 || ny >= len(m) || nx < 0 || nx >= len(m[0]) {
					continue
				}

				if m[ny][nx] == '#' {
					continue
				}

				if ret[Coord{x, y}] == nil {
					ret[Coord{x, y}] = make(map[Coord]int)
				}

				ret[Coord{x, y}][Coord{nx, ny}] = 1
			}
		}
	}

	return m, ret
}

func max(nums []int) (ret int) {
	for _, n := range nums {
		if n > ret {
			ret = n
		}
	}
	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	m, graph = parse(input)

	lens := graph.DFS(Coord{1, 0}, m.End(), map[Coord]bool{{1, 0}: true}, true)

	return max(lens)
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	m, graph = parse(input)

	graph.Contract()

	lens := graph.DFS(Coord{1, 0}, m.End(), map[Coord]bool{}, false)

	return max(lens)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
