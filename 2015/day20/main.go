package main

import (
	"fmt"
	"os"
)

func parse(lines string) (ret int) {
	fmt.Sscanf(lines, "%d", &ret)
	return ret
}

// brute, you?
func partOne(lines string) int {
	num := parse(lines)

	houses := make([]int, num)

	for elf := 1; elf < num; elf++ {
		for house := elf; house < num; house += elf {
			p := houses[house] + elf*10
			houses[house] = p
		}
	}

	for house, presents := range houses {
		if presents >= num {
			return house
		}
	}
	return 0
}

func partTwo(lines string) (ret int) {
	num := parse(lines)

	houses := make([]int, num)

	for elf := 1; elf < num; elf++ {
		i := 0
		for house := elf; house < num; house += elf {
			p := houses[house] + elf*11
			houses[house] = p
			i += 1
			if i > 50 {
				break
			}
		}
	}

	for house, presents := range houses {
		if presents >= num {
			return house
		}
	}

	return 0
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
