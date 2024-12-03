package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(input string) int {
	var ret int

	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			ret += int(input[i] - '0')
		}
	}

	return ret
}

func part2(input string) int {
	var ret int

	for i := 0; i < len(input)/2; i++ {
		if input[i] == input[i+len(input)/2] {
			ret += int(input[i]-'0') * 2
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(strings.TrimSpace(string(data))))
	fmt.Println("part2", part2(strings.TrimSpace(string(data))))
}
