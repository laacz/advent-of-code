package main

import (
	"fmt"
	"os"
	"strings"
)

// min returns the int from the string, falling back to 0
func getInt(s string) (ret int) {
	fmt.Sscanf(s, "%d", &ret)
	return ret
}

// parse parses the input into required form
func parse(input string) (ret []int) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ret = append(ret, getInt(line))
	}

	return ret
}

// sum retruns the sum of all numbers in nums
func sum(nums []int) (ret int) {
	for _, num := range nums {
		ret += num
	}
	return ret
}

// mul returns the product of all numbers in nums
func mul(nums []int) (ret int) {
	ret = 1
	for _, num := range nums {
		ret *= num
	}
	return ret
}

// combinations returns all combinations of length num from arr, built recursively
func combinations(arr []int, num int) [][]int {
	ret := [][]int{}
	if num == 1 {
		for _, a := range arr {
			ret = append(ret, []int{a})
		}
		return ret
	}

	for i := 0; i < len(arr)-num+1; i++ {
		for _, comb := range combinations(arr[i+1:], num-1) {
			ret = append(ret, append(comb, arr[i]))
		}
	}
	return ret
}

func findMinQE(nums []int, w int) (ret int) {
	length := 2
	ret = int(^uint(0) >> 1)
	done := false
	// while we've not found set that sum to w with a minimum number of packages
	for !done {
		combos := combinations(nums, length)
		for _, combo := range combos {
			if sum(combo) == w {
				ret = min(mul(combo), ret)
				// we've found the minimum number of packages that sum to w
				// set the flag, get all combos and then break the main loop
				done = true
			}
		}
		length += 1
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle
func partOne(lines string) (ret int) {
	nums := parse(lines)

	return findMinQE(nums, sum(nums)/3)
}

// partTwo returns the answer to part two of this day's puzzle
func partTwo(lines string) (ret int) {
	nums := parse(lines)

	return findMinQE(nums, sum(nums)/4)

}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
