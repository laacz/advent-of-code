package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	char    byte
	options [][]int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", part1(string(data)))
	fmt.Println("Part 2:", part2(string(data)))
}

func parse(input string) (map[int]Rule, []string) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := make(map[int]Rule)

	for _, line := range strings.Split(parts[0], "\n") {
		colonIdx := strings.Index(line, ":")
		id, _ := strconv.Atoi(line[:colonIdx])
		rest := strings.TrimSpace(line[colonIdx+1:])

		if rest[0] == '"' {
			rules[id] = Rule{char: rest[1]}
		} else {
			var options [][]int
			for _, opt := range strings.Split(rest, " | ") {
				var nums []int
				for _, n := range strings.Fields(opt) {
					num, _ := strconv.Atoi(n)
					nums = append(nums, num)
				}
				options = append(options, nums)
			}
			rules[id] = Rule{options: options}
		}
	}

	messages := strings.Split(parts[1], "\n")
	return rules, messages
}

func match(rules map[int]Rule, ruleID int, msg string, pos int) []int {
	if pos >= len(msg) {
		return nil
	}

	rule := rules[ruleID]

	if rule.char != 0 {
		if msg[pos] == rule.char {
			return []int{pos + 1}
		}
		return nil
	}

	var results []int
	for _, option := range rule.options {
		positions := []int{pos}
		for _, subRule := range option {
			var newPositions []int
			for _, p := range positions {
				newPositions = append(newPositions, match(rules, subRule, msg, p)...)
			}
			positions = newPositions
			if len(positions) == 0 {
				break
			}
		}
		results = append(results, positions...)
	}

	return results
}

func matches(rules map[int]Rule, msg string) bool {
	positions := match(rules, 0, msg, 0)
	for _, p := range positions {
		if p == len(msg) {
			return true
		}
	}
	return false
}

func part1(input string) int {
	rules, messages := parse(input)
	count := 0
	for _, msg := range messages {
		if matches(rules, msg) {
			count++
		}
	}
	return count
}

func part2(input string) int {
	rules, messages := parse(input)

	rules[8] = Rule{options: [][]int{{42}, {42, 8}}}
	rules[11] = Rule{options: [][]int{{42, 31}, {42, 11, 31}}}

	count := 0
	for _, msg := range messages {
		if matches(rules, msg) {
			count++
		}
	}
	return count
}
