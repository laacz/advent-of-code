package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(input string) [][]string {
	ret := [][]string{}

	for _, group := range strings.Split(strings.TrimSpace(input), "\n\n") {
		ret = append(ret, strings.Split(group, "\n"))
	}

	return ret
}
func part1(input [][]string) int {
	ret := 0

	for _, group := range input {
		answers := map[rune]int{}
		for _, answer := range group {
			for _, char := range answer {
				answers[char]++
			}
		}
		ret += len(answers)
	}

	return ret
}

func part2(input [][]string) int {
	ret := 0

	for _, group := range input {
		answers := map[rune]int{}
		for _, answer := range group {
			for _, char := range answer {
				answers[char]++
			}
		}
		for _, cnt := range answers {
			if cnt == len(group) {
				ret++
			}
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
