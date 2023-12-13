package main

import (
	"fmt"
	"os"
	"strings"
)

func isDistinct(s string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) > 1 {
			return false
		}
	}
	return true
}

func partOne(input string) int {
	for i := 4; i < len(input); i++ {
		if isDistinct(input[i-4 : i]) {
			return i
		}
	}

	return 0
}

func partTwo(input string) int {
	for i := 14; i < len(input); i++ {
		if isDistinct(input[i-14 : i]) {
			return i
		}
	}

	return 0
}
func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
