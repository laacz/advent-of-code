package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func caloriesPerElf(input string) []int {
	elves := make([]int, 0)
	sum := 0
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || i == len(lines)-1 {
			elves = append(elves, sum)
			sum = 0
			continue
		}
		i, _ := strconv.Atoi(line)
		sum += i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	return elves
}

func partOne(elves []int) int {
	return elves[0]
}

func partTwo(elves []int) int {
	return elves[0] + elves[1] + elves[2]
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	elves := caloriesPerElf(string(data))
	fmt.Printf("Part one: %d\n", partOne(elves))
	fmt.Printf("Part two: %d\n", partTwo(elves))
}
