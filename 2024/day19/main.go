package main

import (
	"fmt"
	"os"
	"strings"
)

type Input struct {
	Patterns []string
	Designs  []string
}

var impossible = map[string]bool{}

func (i *Input) IsDesignPossible(design string) bool {
	if _, ok := impossible[design]; ok {
		return false
	}

	if len(design) == 0 {
		return true
	}

	for _, p := range i.Patterns {
		if strings.HasPrefix(design, p) {
			if i.IsDesignPossible(design[len(p):]) {
				return true
			}
		}
	}

	impossible[design] = true
	return false
}

var memo = map[string]int{}

func (i *Input) CountPatternCombinations(design string) int {
	if v, ok := memo[design]; ok {
		return v
	}

	if len(design) == 0 {
		return 1
	}

	count := 0
	for _, p := range i.Patterns {
		if strings.HasPrefix(design, p) {
			count += i.CountPatternCombinations(design[len(p):])
		}
	}

	memo[design] = count

	return count
}

func parseInput(input string) Input {
	var ret Input

	parts := strings.Split(input, "\n\n")

	for _, p := range strings.Split(parts[0], ", ") {
		ret.Patterns = append(ret.Patterns, p)
	}

	for _, p := range strings.Split(parts[1], "\n") {
		ret.Designs = append(ret.Designs, p)
	}

	return ret
}

func part1(i Input) int {
	var ret int

	for _, d := range i.Designs {
		if i.IsDesignPossible(d) {
			ret++
		}
	}

	return ret
}

func part2(i Input) int {
	var ret int

	for _, d := range i.Designs {
		ret += i.CountPatternCombinations(d)
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
