package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := part1(parse(string(data)))
	fmt.Println("Part 1:", result)

	result2 := part2(parse(string(data)))
	fmt.Println("Part 2:", result2)
}

func parse(input string) []int {
	lines := strings.Split(input, "\n")
	ret := []int{}

	for _, line := range lines {
		if i, err := strconv.Atoi(line); err == nil {
			ret = append(ret, i)
		}
	}

	return ret
}

func part1(input []int) int {
	for i, num := range	input {
		for _, num2 := range input[i+1:] {
			if num+num2 == 2020 {
				return num * num2
			}
		}
	}

	return 0
}


func part2(input []int) int {
	for i, num := range	input {
		for j, num2 := range input[i+1:] {
			for _, num3 := range input[j+1:] {
				if num+num2+num3 == 2020 {
					return num * num2 * num3
				}
			}
		}
	}

	return 0
}
