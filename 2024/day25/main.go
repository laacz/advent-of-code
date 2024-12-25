package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) ([][]int, [][]int) {
	var locks, keys [][]int

	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, block := range blocks {
		code := [5]int{-1, -1, -1, -1, -1}
		lock := false

		for j, line := range strings.Split(block, "\n") {
			if j == 0 && line == "#####" {
				lock = true
			}
			for i, c := range line {
				if c == '#' {
					code[i] += 1
				}
			}
		}

		if lock {
			locks = append(locks, code[:])
		} else {
			keys = append(keys, code[:])
		}
	}

	return locks, keys
}

func part1(locks, keys [][]int) int {
	var ret int

	for _, key := range keys {
		for _, lock := range locks {
			match := true
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] >= 6 {
					match = false
					break
				}
			}
			if match {
				ret += 1
			}
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	locks, keys := parseInput(string(input))
	fmt.Println("part1", part1(locks, keys))
}
