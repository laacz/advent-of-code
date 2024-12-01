package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func part1(list1, list2 []int) int {
	var ret int

	sort.Ints(list1)
	sort.Ints(list2)

	for k, v1 := range list1 {
		v2 := list2[k]
		ret += int(math.Abs(float64(v1 - v2)))
	}

	return ret
}

func part2(list1, list2 []int) int {
	var ret int
	counts := make(map[int]int)

	for _, v := range list2 {
		counts[v]++
	}

	for _, v := range list1 {
		ret += counts[v] * v
	}

	return ret
}

func main() {
	var list1, list2 []int
	var v1, v2 int

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		fmt.Sscanf(line, "%d %d", &v1, &v2)

		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}

	fmt.Println("part1", part1(list1, list2))
	fmt.Println("part2", part2(list1, list2))
}
