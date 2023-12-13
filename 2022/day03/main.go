package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(input string) int {
	score := 0

	for _, rs := range strings.Split(input, "\n") {
		r1 := rs[:len(rs)/2]
		r2 := rs[len(rs)/2:]

		for _, c1 := range r1 {
			if strings.Contains(r2, string(c1)) {
				pr := c1 - 'a' + 1
				if pr < 0 {
					pr = c1 - 'A' + 27
				}
				score += int(pr)

				break
			}
		}

	}

	return score
}

func partTwo(input string) int {
	score := 0
	rs := strings.Split(input, "\n")

	for i := 0; i < len(rs); i += 3 {
		for _, c := range rs[i] {
			if strings.Contains(rs[i+1], string(c)) && strings.Contains(rs[i+2], string(c)) {
				pr := c - 'a' + 1
				if pr < 0 {
					pr = c - 'A' + 27
				}
				score += int(pr)
				break
			}
		}

	}
	return score
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
