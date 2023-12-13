package main

import (
	"fmt"
	"os"
	"strings"
)

func isNice(line string) bool {
	vowels := 0
	double := false
	prev := rune(0)
	for _, c := range line {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
		if c == prev {
			double = true
		}
		prev = c
	}

	if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
		return false
	}

	if vowels < 3 || !double {
		return false
	}

	return true
}

func isNewNice(line string) bool {
	pair := false
	repeated := false

	for i := 0; i < len(line)-2; i++ {
		if strings.Contains(line[i+2:], line[i:i+2]) {
			pair = true
		}
		if line[i] == line[i+2] {
			repeated = true
		}
	}

	return pair && repeated
}

func partOne(lines string) int {
	ret := 0

	for _, l := range strings.Split(lines, "\n") {
		if isNice(l) {
			ret++
		}
	}

	return ret
}

func partTwo(lines string) int {
	ret := 0

	for _, l := range strings.Split(lines, "\n") {
		if isNewNice(l) {
			ret++
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
