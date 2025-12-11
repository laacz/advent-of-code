package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func countTreesEncounteredWhileTraversingTheSlopeDownwards(input []string, right, down int) int {
	ret, x, y := 0, 0, 0

	for {
		if y >= len(input) {
			break
		}
		if input[y][x%len(input[y])] == '#' {
			ret++
		}
		x += right
		y += down
	}

	return ret
}

func part1(input []string) int {
	return countTreesEncounteredWhileTraversingTheSlopeDownwards(input, 3, 1)
}

func part2(input []string) int {
	ret := 1

	for _, slope := range [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		ret *= countTreesEncounteredWhileTraversingTheSlopeDownwards(input, slope[0], slope[1])
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
