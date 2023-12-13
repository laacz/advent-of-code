package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Graph struct {
	edges map[string][]Edge
	nodes map[string]bool
}

type Edge struct {
	from     string
	to       string
	distance int
}

func parse(data string) Graph {
	ret := Graph{
		edges: make(map[string][]Edge),
		nodes: make(map[string]bool),
	}

	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		var from, to string
		var distance int

		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)

		ret.nodes[from] = true
		ret.nodes[to] = true

		ret.edges[from] = append(ret.edges[from], Edge{
			from:     from,
			to:       to,
			distance: distance,
		})

		ret.edges[to] = append(ret.edges[to], Edge{
			from:     to,
			to:       from,
			distance: distance,
		})
	}

	return ret
}

func (g *Graph) Distance(path []string) int {
	totalDistance := 0
	for i := 0; i < len(path)-1; i++ {
		for _, edge := range g.edges[path[i]] {
			if edge.to == path[i+1] {
				totalDistance += edge.distance
				break
			}
		}
	}
	return totalDistance
}

// Generate all permutations of a slice of strings
func permute(values []string, start int, result *[][]string) {
	if start == len(values)-1 {
		permutation := make([]string, len(values))
		copy(permutation, values)
		*result = append(*result, permutation)
		return
	}

	for i := start; i < len(values); i++ {
		values[start], values[i] = values[i], values[start]
		permute(values, start+1, result)
		values[start], values[i] = values[i], values[start]
	}
}

func (g *Graph) ShortestPathTSP() []string {
	nodes := make([]string, 0, len(g.nodes))
	for node := range g.nodes {
		nodes = append(nodes, node)
	}

	allPermutations := [][]string{}
	permute(nodes, 0, &allPermutations)

	shortestPath := []string{}
	minDistance := math.MaxInt

	for _, path := range allPermutations {
		distance := g.Distance(path)
		if distance < minDistance {
			minDistance = distance
			shortestPath = path
		}
	}

	return shortestPath
}

func (g *Graph) LongestPath() []string {
	nodes := make([]string, 0, len(g.nodes))
	for node := range g.nodes {
		nodes = append(nodes, node)
	}

	allPermutations := [][]string{}
	permute(nodes, 0, &allPermutations)

	longestPath := []string{}
	maxDistance := 0

	for _, path := range allPermutations {
		distance := g.Distance(path)
		if distance > maxDistance {
			maxDistance = distance
			longestPath = path
		}
	}

	return longestPath
}

func partOne(lines string) (ret int) {
	g := parse(lines)
	ret = g.Distance(g.ShortestPathTSP())

	return ret
}

func partTwo(lines string) (ret int) {
	g := parse(lines)
	ret = g.Distance(g.LongestPath())

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
