package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(input string) []int {
	ret := []int{}

	for line := range strings.Lines(strings.TrimSpace(input)) {
		num, _ := strconv.Atoi(strings.TrimSpace(line))
		ret = append(ret, num)
	}

	return ret
}

func part1(input []int) int {
	adapters := make([]int, len(input))
	copy(adapters, input)
	slices.Sort(adapters)

	diffs := map[int]int{}
	prev := 0

	for _, adapter := range adapters {
		diffs[adapter-prev]++
		prev = adapter
	}

	diffs[3]++

	return diffs[1] * diffs[3]
}

func part2(input []int) int {
	adapters := make([]int, len(input))
	copy(adapters, input)
	slices.Sort(adapters)

	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	ways := map[int]int{}
	ways[0] = 1

	for i := 1; i < len(adapters); i++ {
		for j := i - 1; j >= 0 && adapters[i]-adapters[j] <= 3; j-- {
			ways[adapters[i]] += ways[adapters[j]]
		}
	}

	return ways[adapters[len(adapters)-1]]
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
