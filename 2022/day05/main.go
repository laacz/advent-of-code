package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) ([][]string, [][]int) {
	parts := strings.Split(input, "\n\n")

	max_len := 0
	for _, stack := range strings.Split(parts[0], "\n") {
		if max_len < len(stack)/4+1 {
			max_len = len(stack)/4 + 1
		}
	}

	s, m := make([][]string, max_len), [][]int{}

	for _, stack := range strings.Split(parts[0], "\n") {
		if !strings.Contains(stack, "[") {
			continue
		}
		pos := 0
		for pos <= len(stack) {
			crate := stack[pos+1]
			if crate != ' ' {
				s[pos/4] = append(s[pos/4], string(crate))
			}
			pos += 4
		}
	}

	for _, move := range strings.Split(parts[1], "\n") {
		amt, f, t := 0, 0, 0
		fmt.Sscanf(move, "move %d from %d to %d", &amt, &f, &t)
		m = append(m, []int{amt, f, t})
	}

	return s, m
}

func partOne(input string) string {
	score := ""

	stacks, moves := parseInput(input)

	for _, move := range moves {
		amt, from, to := move[0], move[1], move[2]

		for i := 0; i < amt; i++ {
			stacks[to-1] = append([]string{stacks[from-1][0]}, stacks[to-1]...)
			stacks[from-1] = stacks[from-1][1:]
		}
	}

	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		score += stack[0]
	}

	return score
}

func partTwo(input string) string {
	score := ""

	stacks, moves := parseInput(input)

	for _, move := range moves {
		amt, from, to := move[0], move[1], move[2]

		for i := amt - 1; i >= 0; i-- {
			stacks[to-1] = append([]string{stacks[from-1][i]}, stacks[to-1]...)
		}

		stacks[from-1] = stacks[from-1][amt:]
	}

	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		score += stack[0]
	}

	return score
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %s\n", partOne(string(data)))
	fmt.Printf("Part two: %s\n", partTwo(string(data)))
}
