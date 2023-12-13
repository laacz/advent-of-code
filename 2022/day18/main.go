package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(input string) int {
	ret := 0
	cubes := make(map[string]bool)

	for _, line := range strings.Split(input, "\n") {
		cubes[line] = true
		ret += 6
	}

	for k := range cubes {
		var x, y, z int
		_, _ = fmt.Sscanf(k, "%d,%d,%d", &x, &y, &z)

		if cubes[fmt.Sprintf("%d,%d,%d", x-1, y, z)] {
			ret--
		}
		if cubes[fmt.Sprintf("%d,%d,%d", x, y-1, z)] {
			ret--
		}
		if cubes[fmt.Sprintf("%d,%d,%d", x, y, z-1)] {
			ret--
		}
		if cubes[fmt.Sprintf("%d,%d,%d", x+1, y, z)] {
			ret--
		}
		if cubes[fmt.Sprintf("%d,%d,%d", x, y+1, z)] {
			ret--
		}
		if cubes[fmt.Sprintf("%d,%d,%d", x, y, z+1)] {
			ret--
		}
	}

	return ret
}

type Cube struct {
	x, y, z int
}

func (c *Cube) neighbors() []Cube {
	var ret []Cube
	for _, c := range []Cube{
		{c.x - 1, c.y, c.z},
		{c.x + 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z - 1},
		{c.x, c.y, c.z + 1},
	} {
		if c.x < -1 || c.y < -1 || c.z < -1 ||
			c.x > 20 || c.y > 20 || c.z > 20 {
			continue
		}
		ret = append(ret, c)
	}

	return ret
}

func partTwo(input string) int {
	ret := 0

	cubes := make(map[Cube]bool)
	visited := make(map[Cube]bool)

	for _, line := range strings.Split(input, "\n") {
		var c Cube
		_, _ = fmt.Sscanf(line, "%d,%d,%d", &c.x, &c.y, &c.z)
		cubes[c] = true
	}

	stack := []Cube{
		{0, 0, 0},
	}

	for len(stack) > 0 {
		s := stack[0]
		stack = stack[1:]

		for _, c := range s.neighbors() {
			if cubes[c] {
				ret += 1
			} else if !visited[c] {
				visited[Cube{c.x, c.y, c.z}] = true
				stack = append(stack, c)
			}
		}
	}

	return ret
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
