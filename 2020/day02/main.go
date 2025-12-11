package main

import (
	"fmt"
	"os"
	"strings"
)

type Password struct {
	Min  int
	Max  int
	Char rune
	Text string
}

func parse(input string) []Password {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var rules []Password
	for _, line := range lines {
		var p Password
		fmt.Sscanf(line, "%d-%d %c: %s", &p.Min, &p.Max, &p.Char, &p.Text)
		rules = append(rules, p)
	}

	return rules
}

func part1(passwords []Password) int {
	ret := 0

	for _, p := range passwords {
		cnt := strings.Count(p.Text, string(p.Char))
		if cnt >= p.Min && cnt <= p.Max {
			ret += 1
		}
	}

	return ret
}

func part2(passwords []Password) int {
	ret := 0

	for _, p := range passwords {
		if (rune(p.Text[p.Min-1]) == p.Char) != (rune(p.Text[p.Max-1]) == p.Char) {
			ret += 1
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
