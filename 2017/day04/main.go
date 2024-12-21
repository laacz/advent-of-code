package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) [][]string {
	var ret [][]string

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		ret = append(ret, strings.Split(line, " "))
	}

	return ret
}

func part1(input [][]string) int {
	var ret int

	for _, line := range input {
		words := make(map[string]bool)
		valid := true
		for _, word := range line {
			if words[word] {
				valid = false
				break
			}
			words[word] = true
		}
		if valid {
			ret++
		}
	}

	return ret
}

func areAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	chars := make(map[rune]int)
	for _, r := range a {
		chars[r]++
	}
	for _, r := range b {
		chars[r]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

func part2(input [][]string) int {
	var ret int
	for _, line := range input {
		words := make(map[string]bool)
		valid := true
	outer:
		for _, word := range line {
			for k := range words {
				if areAnagrams(k, word) {
					valid = false
					break outer
				}
			}
			words[word] = true
		}

		if valid {
			ret++
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(data))))
	fmt.Println("part2", part2(parseInput(string(data))))
}
