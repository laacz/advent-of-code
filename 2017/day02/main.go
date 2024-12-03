package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(input string) int {
	var ret int

	for _, line := range strings.Split(input, "\n") {
		var min, max int
		max = 0
		min = math.MaxInt

		slice := strings.Fields(line)
		for v := range slice {
			num, _ := strconv.Atoi(slice[v])
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		ret += max - min
	}

	return ret
}

func part2(input string) int {
	var ret int

	for _, line := range strings.Split(input, "\n") {
		slice := strings.Fields(line)

		for i := 0; i < len(slice); i++ {
			a, _ := strconv.Atoi(slice[i])
			for j := 0; j < len(slice); j++ {
				if i == j {
					continue
				}
				b, _ := strconv.Atoi(slice[j])
				if a%b == 0 {
					ret += a / b
				}
			}
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(strings.TrimSpace(string(data))))
	fmt.Println("part2", part2(strings.TrimSpace(string(data))))
}
