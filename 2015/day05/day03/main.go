package main

import (
	"fmt"
	"os"
)

func move(curr *[2]int, c rune) {
	switch c {
	case '^':
		curr[1]++
	case 'v':
		curr[1]--
	case '<':
		curr[0]--
	case '>':
		curr[0]++
	}
}

func partOne(lines string) int {
	houses := make(map[[2]int]int)
	curr := [2]int{0, 0}
	houses[curr]++

	for _, c := range lines {
		move(&curr, c)
		houses[curr]++
	}

	return len(houses)
}

func partTwo(lines string) int {
	houses := make(map[[2]int]int)
	santa := [2]int{0, 0}
	robo := [2]int{0, 0}
	houses[santa]++
	houses[robo]++

	i := 0
	for _, c := range lines {
		i += 1
		if i%2 == 0 {
			move(&robo, c)
			houses[robo]++
		} else {
			move(&santa, c)
			houses[santa]++
		}
	}

	return len(houses)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
