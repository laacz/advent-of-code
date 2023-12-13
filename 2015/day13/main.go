package main

import (
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	edges map[string]map[string]int
}

func parse(line string) Graph {
	ret := Graph{
		edges: make(map[string]map[string]int),
	}
	for _, l := range strings.Split(line, "\n") {
		if l == "" {
			continue
		}

		var from, to, sign string
		var units int

		fmt.Sscanf(l, "%s would %s %d happiness units by sitting next to %s.", &from, &sign, &units, &to)

		to = strings.TrimSuffix(to, ".")

		if sign == "lose" {
			units = -units
		}

		if _, ok := ret.edges[from]; !ok {
			ret.edges[from] = make(map[string]int)
		}

		ret.edges[from][to] = units
	}

	return ret
}

func (g Graph) String() (ret string) {
	for from, edges := range g.edges {
		for to, distance := range edges {
			ret += fmt.Sprintf("%s -> %s = %d\n", from, to, distance)
		}
	}

	return ret
}

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

func partOne(line string) (ret int) {
	g := parse(line)

	arr := []string{}
	for from := range g.edges {
		arr = append(arr, from)
	}

	perms := [][]string{}
	permute(arr, 0, &perms)

	for _, perm := range perms {
		distance := 0

		dbg := ""

		for i := 0; i < len(perm)-1; i++ {
			dbg += fmt.Sprint(perm[i], " ")

			a := g.edges[perm[i]][perm[i+1]]
			b := g.edges[perm[i+1]][perm[i]]

			distance += a + b

			dbg += fmt.Sprint(a+b, " ")
		}

		distance += g.edges[perm[0]][perm[len(perm)-1]]
		distance += g.edges[perm[len(perm)-1]][perm[0]]

		if distance > ret {
			ret = distance
		}
		dbg += fmt.Sprint(perm, ret)
	}

	return ret
}

func partTwo(line string) (ret int) {
	g := parse(line)

	arr := []string{"laacz"}
	for from := range g.edges {
		arr = append(arr, from)
	}

	perms := [][]string{}
	permute(arr, 0, &perms)

	for _, perm := range perms {
		distance := 0

		dbg := ""

		for i := 0; i < len(perm)-1; i++ {
			dbg += fmt.Sprint(perm[i], " ")

			a := g.edges[perm[i]][perm[i+1]]
			b := g.edges[perm[i+1]][perm[i]]

			distance += a + b

			dbg += fmt.Sprint(a+b, " ")
		}

		distance += g.edges[perm[0]][perm[len(perm)-1]]
		distance += g.edges[perm[len(perm)-1]][perm[0]]

		if distance > ret {
			ret = distance
		}
		dbg += fmt.Sprint(perm, ret)
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
