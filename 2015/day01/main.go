package main

import (
	"fmt"
	"os"
)

func partOne(lines string) int {
	sum := 0

	for _, c := range lines {
		if c == '(' {
			sum++
		} else if c == ')' {
			sum--
		}
	}

	return sum
}

func partTwo(lines string) int {
	pos := 0
	for i, c := range lines {
		if c == '(' {
			pos++
		} else if c == ')' {
			pos--
		}

		if pos == -1 {
			return i + 1
		}
	}

	return 0
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
