package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Network map[string]*Connections
type Connections []string

func (c *Connections) Contains(s string) bool {
	for _, v := range *c {
		if v == s {
			return true
		}
	}
	return false
}

func parseInput(input string) Network {
	ret := make(Network)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, "-")

		if _, ok := ret[parts[0]]; !ok {
			ret[parts[0]] = &Connections{}
		}

		if _, ok := ret[parts[1]]; !ok {
			ret[parts[1]] = &Connections{}
		}

		*ret[parts[0]] = append(*ret[parts[0]], parts[1])
		*ret[parts[1]] = append(*ret[parts[1]], parts[0])

	}
	return ret
}

func part1(network Network) int {
	triplets := make(map[string]bool)

	for a, connections := range network {
		if !strings.HasPrefix(a, "t") {
			continue
		}
		for _, b := range *connections {
			for _, c := range *network[b] {
				if (*network[c]).Contains(a) {
					key := []string{a, b, c}
					sort.Strings(key)
					triplets[strings.Join(key, "-")] = true
				}
			}
		}
	}

	return len(triplets)
}

func findMaximalClique(network Network, nodes []string, resultSet []string, excluded []string, largestClique *[]string) []string {
	if len(nodes) == 0 && len(excluded) == 0 {
		if len(resultSet) > len(*largestClique) {
			*largestClique = make([]string, len(resultSet))
			copy(*largestClique, resultSet)
		}
	}

	nodesCopy := make([]string, len(nodes))
	copy(nodesCopy, nodes)

	for _, v := range nodesCopy {
		adj := *network[v]

		var intersectNodes []string
		for _, n := range nodes {
			for _, a := range adj {
				if n == a {
					intersectNodes = append(intersectNodes, n)
					break
				}
			}
		}

		var intersectExcluded []string
		for _, e := range excluded {
			for _, a := range adj {
				if e == a {
					intersectExcluded = append(intersectExcluded, e)
					break
				}
			}
		}

		findMaximalClique(
			network,
			intersectNodes,
			append(resultSet, v),
			intersectExcluded,
			largestClique,
		)

		var newNodes []string
		for _, n := range nodes {
			if n != v {
				newNodes = append(newNodes, n)
			}
		}
		nodes = newNodes

		excluded = append(excluded, v)
	}

	return *largestClique
}

func part2(network Network) string {
	var ret []string

	var nodes []string
	for k := range network {
		nodes = append(nodes, k)
	}

	findMaximalClique(network, nodes, []string{}, []string{}, &ret)

	sort.Strings(ret)

	return strings.Join(ret, ",")
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
