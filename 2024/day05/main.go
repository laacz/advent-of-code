package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) (Rules, []Update) {
	var rules Rules
	parts := strings.Split(input, "\n\n")
	for _, rule := range strings.Split(parts[0], "\n") {
		parts := strings.Split(rule, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		rules = append(rules, []int{first, second})
	}

	var updates []Update
	for _, line := range strings.Split(parts[1], "\n") {
		var update []int
		for _, num := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(num)
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

type Rules [][]int

func (r Rules) Match(first, second int) bool {
	for _, rule := range r {
		if rule[0] == second && rule[1] == first {
			return false
		}
	}

	return true
}

type Update []int

func (u *Update) Sorted(rules Rules) bool {
	for i, a := range *u {
		for _, b := range (*u)[i+1:] {
			if !rules.Match(a, b) {
				return false
			}
		}
	}

	return true
}

func (u *Update) Fix(rules Rules) {
	for i, a := range *u {
		for j, b := range (*u)[i+1:] {
			if !rules.Match(a, b) {
				(*u)[i], (*u)[i+j+1] = (*u)[i+j+1], (*u)[i]
			}
		}
	}
}

func part1(input string) int {
	var ret int

	rules, updates := parseInput(input)

	for _, update := range updates {
		if update.Sorted(rules) {
			ret += update[len(update)/2]
		}
	}

	return ret
}

func part2(input string) int {
	var ret int

	rules, updates := parseInput(input)

	for _, update := range updates {
		if !update.Sorted(rules) {
			for !update.Sorted(rules) {
				update.Fix(rules)
			}
			ret += update[len(update)/2]
		}
	}

	return ret
}

func main() {

	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
