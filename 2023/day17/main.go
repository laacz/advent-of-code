package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(input string) (ret [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, c := range line {
			var num int
			fmt.Sscanf(string(c), "%d", &num)
			row = append(row, num)
		}
		ret = append(ret, row)
	}
	return ret
}

type Cell struct {
	x, y   int
	dir    [2]int
	conseq int
}

// minHeatLoss shortest weighted path according to Dijkstra. Implemented wihtout
// a priority queue. That's why it's nost particulary fast.
func minHeatLoss(grid [][]int, minconseq, maxconseq int) (ret int) {
	// add first v to Q, two as a matter of fact - each entering from different
	// direction
	queue := []Cell{
		{0, 0, [2]int{1, 0}, 0},
		{0, 0, [2]int{0, 1}, 0},
	}
	// dist[source] ← 0
	dist := map[Cell]int{{0, 0, [2]int{0, 0}, 0}: 0}
	// dist is inf
	ret = int(^uint(0) >> 1)

	for len(queue) > 0 {
		// u ← vertex in Q with min dist[u]
		c := queue[0]
		// remove u from Q
		queue = queue[1:]

		// have we reached the end?
		if c.x == len(grid)-1 && c.y == len(grid[0])-1 {
			ret = min(ret, dist[c])
		}

		// for each neighbor v of u still in Q. Note that directions are only forward,
		// left and right. We cannot go backwards.
		for _, dir := range [][2]int{c.dir, {c.dir[1], -c.dir[0]}, {-c.dir[1], c.dir[0]}} {
			// next cell's coordinates
			x, y := c.x+dir[0], c.y+dir[1]

			// are we out of bounds?
			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
				continue
			}

			// alt ← dist[u] + Graph.Edges(u, v)
			totalheatloss := dist[c] + grid[x][y]

			// consequetive cells in the same direction
			conseq := 1
			if c.dir == dir {
				conseq = c.conseq + 1
			}

			// create next cell
			next := Cell{x, y, dir, conseq}

			// Check if it's already measured. Note that we're declaring uniqueness
			// by the coordinates, direction AND number of consequetive cells in the
			// same direction. It's important!
			cc := dist[next]

			// if alt < dist[v]:
			if cc == 0 || cc > totalheatloss {
				if (dir == c.dir && c.conseq < maxconseq) ||
					(dir != c.dir && c.conseq >= minconseq) {
					queue = append(queue, next)

					// dist[v] ← alt
					dist[next] = totalheatloss
				}
			}
		}
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	g := parse(input)
	ret = minHeatLoss(g, 0, 3)
	return ret
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	g := parse(input)
	ret = minHeatLoss(g, 4, 10)
	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
