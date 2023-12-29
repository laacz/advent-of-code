package main

import (
	"fmt"
	"os"
	"strings"
)

type Graph map[string][]string

func parse(input string, cut [][]string) (ret Graph) {
	ret = make(Graph)
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		from := strings.TrimSpace(parts[0])
		to := strings.Split(strings.TrimSpace(parts[1]), " ")

		if _, ok := ret[from]; !ok {
			ret[from] = []string{}
		}

		for _, t := range to {
			if _, ok := ret[t]; !ok {
				ret[t] = []string{}
			}

			skip := false
			for _, w := range cut {
				if (from == w[0] && t == w[1]) || (from == w[1] && t == w[0]) {
					skip = true
					break
				}
			}

			if !skip {
				ret[t] = append(ret[t], from)
				ret[from] = append(ret[from], t)
			}
		}

	}
	return ret
}

func (g *Graph) findPath(from, to string) (ret []string) {
	visited := make(map[string]bool)
	queue := []string{from}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == to {
			ret = append(ret, current) // Append the current node to ret
			return ret
		}
		visited[current] = true
		for _, n := range (*g)[current] {
			if _, ok := visited[n]; !ok {
				queue = append(queue, n)
			}
		}
		ret = append(ret, current) // Append the current node to ret
	}
	if ret[len(ret)-1] != to {
		return []string{}
	}
	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string, cut [][]string) (ret int) {
	g := parse(input, cut)

	names := []string{}
	for k := range g {
		names = append(names, k)
	}

	groups := []map[string]bool{
		make(map[string]bool),
		make(map[string]bool),
	}
	for i, start := range cut[0] {
		groups[i][start] = true
		for end := range g {
			if start == end {
				continue
			}
			path := g.findPath(start, end)
			if len(path) > 0 {
				groups[i][end] = true
			}
		}

	}

	return len(groups[0]) * len(groups[1])
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data), [][]string{
		{"tpn", "gxv"},
		{"zcj", "rtt"},
		{"hxq", "txl"},
	}))

}
