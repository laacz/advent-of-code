package main

import (
	"fmt"
	"os"
	"strings"
)

func charToElevation(name byte) int {
	if name == 'S' {
		name = 'a'
	}
	if name == 'E' {
		name = 'z'
	}
	return int(name - 'a')
}

func parseInputToGraph(input string) ([]string, string, string, map[string][]string) {
	graph := make(map[string][]string)
	lines := strings.Split(input, "\n")
	var start string
	var end string
	var starts []string
	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, elevation := range line {
			name := fmt.Sprintf("%d,%d", y, x)
			graph[name] = []string{}
			if elevation == 'S' {
				start = name
			} else if elevation == 'E' {
				end = name
			}
			elevation := charToElevation(byte(elevation))
			if elevation == 0 {
				starts = append(starts, name)
			}
			for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				yy := y + dir[0]
				xx := x + dir[1]
				if xx >= 0 && yy >= 0 && xx < len(line) && yy < len(lines) {
					neighbour := fmt.Sprintf("%d,%d", yy, xx)
					if charToElevation(lines[yy][xx])-elevation < 2 {
						graph[name] = append(graph[name], neighbour)
					}
				}
			}
		}
	}
	return starts, start, end, graph
}

func shortestPath(graph map[string][]string, start string, end string) int {
	distances := make(map[string]int)
	seen := []string{start}
	last := 0
loop:
	for i, point := range seen {
		if i < last {
			continue
		}
		last = i
		steps := distances[point]
		for _, neighbour := range graph[point] {
			_, ok := distances[neighbour]
			if !ok || distances[neighbour] > steps+1 {
				distances[neighbour] = steps + 1
				seen = append(seen, neighbour)
				goto loop
			}
		}
	}
	return distances[end]
}

func partOne(input string) int {
	_, start, end, graph := parseInputToGraph(input)
	return shortestPath(graph, start, end)
}

func partTwo(input string) int {
	starts, _, end, graph := parseInputToGraph(input)
	shortest := 9999
	for _, start := range starts {
		res := shortestPath(graph, start, end)
		if res < shortest && res > 0 {
			shortest = res
		}
	}
	return shortest
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
