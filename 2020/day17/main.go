package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [4]int

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", part1(string(data)))
	fmt.Println("Part 2:", part2(string(data)))
}

func parse(input string) map[Coord]bool {
	active := make(map[Coord]bool)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				var c Coord
				c[0], c[1] = x, y
				active[c] = true
			}
		}
	}

	return active
}

func getNeighbors(c Coord, dims int) []Coord {
	var neighbors []Coord
	offsets := make([]int, dims)
	for i := range offsets {
		offsets[i] = -1
	}

	for {
		isOrigin := true
		for i := 0; i < dims; i++ {
			if offsets[i] != 0 {
				isOrigin = false
				break
			}
		}

		if !isOrigin {
			var neighbor Coord
			for i := 0; i < dims; i++ {
				neighbor[i] = c[i] + offsets[i]
			}
			neighbors = append(neighbors, neighbor)
		}

		i := 0
		for i < dims {
			offsets[i]++
			if offsets[i] <= 1 {
				break
			}
			offsets[i] = -1
			i++
		}
		if i == dims {
			break
		}
	}

	return neighbors
}

func simulate(active map[Coord]bool, dims int) map[Coord]bool {
	candidates := make(map[Coord]bool)

	for c := range active {
		candidates[c] = true
		for _, n := range getNeighbors(c, dims) {
			candidates[n] = true
		}
	}

	newActive := make(map[Coord]bool)
	for c := range candidates {
		neighborCount := 0
		for _, n := range getNeighbors(c, dims) {
			if active[n] {
				neighborCount++
			}
		}

		if active[c] {
			if neighborCount == 2 || neighborCount == 3 {
				newActive[c] = true
			}
		} else {
			if neighborCount == 3 {
				newActive[c] = true
			}
		}
	}

	return newActive
}

func part1(input string) int {
	active := parse(input)

	for cycle := 0; cycle < 6; cycle++ {
		active = simulate(active, 3)
	}

	return len(active)
}

func part2(input string) int {
	active := parse(input)

	for cycle := 0; cycle < 6; cycle++ {
		active = simulate(active, 4)
	}

	return len(active)
}
