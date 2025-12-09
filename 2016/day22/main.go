package main

import (
	"fmt"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

type Coord [2]int

type Disk struct {
	Used  int
	Avail int
}

type Map struct {
	m map[Coord]Disk
}

func parseInput(input string) Map {
	ret := Map{
		m: make(map[Coord]Disk),
	}

	for _, line := range strings.Split(input, "\n") {
		var x, y, size, used, avail int
		fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT", &x, &y, &size, &used, &avail)
		if size > 0 {
			ret.m[Coord{x, y}] = Disk{used, avail}
		}
	}

	return ret
}

func partOne(input string) int {
	var ret int
	m := parseInput(input)
	// fmt.Println(m)
	for coordA, diskA := range m.m {
		for coordB, diskB := range m.m {
			if coordA == coordB || diskA.Used == 0 {
				continue
			}
			if diskA.Used <= diskB.Avail {
				ret++
			}
		}
	}
	return ret
}

func partTwo(input string) int {
	m := parseInput(input)

	var empty Coord
	var maxX int
	for c, d := range m.m {
		if d.Used == 0 {
			empty = c
		}
		if c[0] > maxX {
			maxX = c[0]
		}
	}

	target := Coord{maxX - 1, 0}
	type state struct {
		pos   Coord
		steps int
	}
	visited := make(map[Coord]bool)
	queue := []state{{empty, 0}}
	visited[empty] = true
	dirs := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var steps int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.pos == target {
			steps = cur.steps
			break
		}

		for _, d := range dirs {
			next := Coord{cur.pos[0] + d[0], cur.pos[1] + d[1]}
			if visited[next] {
				continue
			}
			disk, ok := m.m[next]
			if !ok || disk.Used > 100 {
				continue
			}
			visited[next] = true
			queue = append(queue, state{next, cur.steps + 1})
		}
	}

	steps += 1 + (maxX-1)*5

	return steps
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
