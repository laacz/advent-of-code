package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	card, _ := strconv.Atoi(lines[0])
	door, _ := strconv.Atoi(lines[1])
	return card, door
}

func part1(cardPub, doorPub int) int {
	value, loops := 1, 0
	for value != cardPub {
		value = (value * 7) % 20201227
		loops++
	}

	value = 1
	for i := 0; i < loops; i++ {
		value = (value * doorPub) % 20201227
	}
	return value
}

func main() {
	data, _ := os.ReadFile("input.txt")
	cardPub, doorPub := parse(string(data))

	fmt.Println("Part 1:", part1(cardPub, doorPub))
}
