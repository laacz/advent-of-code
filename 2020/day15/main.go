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

	numbers := parse(string(data))

	result := part1(numbers)
	fmt.Println("Part 1:", result)

	result2 := part2(numbers)
	fmt.Println("Part 2:", result2)
}

func parse(input string) []int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")
	numbers := make([]int, len(parts))
	for i, p := range parts {
		numbers[i], _ = strconv.Atoi(p)
	}
	return numbers
}

func play(starting []int, targetTurn int) int {
	lastSpoken := make([]int32, targetTurn)

	for i := 0; i < len(starting)-1; i++ {
		lastSpoken[starting[i]] = int32(i + 1)
	}

	lastNum := int32(starting[len(starting)-1])
	for turn := int32(len(starting) + 1); turn <= int32(targetTurn); turn++ {
		prev := lastSpoken[lastNum]
		lastSpoken[lastNum] = turn - 1
		if prev == 0 {
			lastNum = 0
		} else {
			lastNum = (turn - 1) - prev
		}
	}

	return int(lastNum)
}

func part1(numbers []int) int {
	return play(numbers, 2020)
}

func part2(numbers []int) int {
	return play(numbers, 30000000)
}
