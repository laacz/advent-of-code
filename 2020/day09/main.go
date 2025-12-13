package main

import (
	"fmt"
	"os"
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

func part1(input []int, preamble int) int {
	for pos, num := range input {
		if pos < preamble {
			continue
		}
		found := false
		candidates := input[pos-preamble : pos]

		for i := range candidates {
			for j := i + 1; j < len(candidates); j++ {
				if candidates[i]+candidates[j] == num {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			return num
		}
	}

	return 0
}

func part2(input []int, preamble int) int {
	target := part1(input, preamble)

	for start := range input {
		for end := start + 1; end <= len(input); end++ {
			subset := input[start:end]
			if target == sum(subset) {
				return min(subset) + max(subset)
			}
		}
	}

	return 0
}

func min(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data)), 25))
	fmt.Println("Part 2:", part2(parse(string(data)), 25))
}
