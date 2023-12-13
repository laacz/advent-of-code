package main

import (
	"fmt"
	"os"
	"strings"
)

func isVisible(lines []string, x, y int) bool {
	deltas := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	height := lines[y][x]
Loop:
	for _, delta := range deltas {
		y1 := y
		x1 := x
		for y1 > 0 && y1 < len(lines)-1 && x1 > 0 && x1 < len(lines[y1])-1 {
			x1 += delta[1]
			y1 += delta[0]
			if lines[y1][x1] >= height {
				continue Loop
			}
		}
		return true
	}
	return false
}

func partOne(input string) int {
	result := 0
	lines := strings.Split(input, "\n")

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			visible := isVisible(lines, x, y)
			if visible {
				result += 1
			}
		}
	}

	return result
}

func partTwo(input string) int {
	result := 0
	lines := strings.Split(input, "\n")

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			visible := map[string]int{
				"up":    1,
				"down":  1,
				"left":  1,
				"right": 1,
			}
			for pos := y - 1; pos > 0 && lines[pos][x] < lines[y][x]; {
				pos--
				visible["up"]++
			}

			for pos := y + 1; pos < len(lines)-1 && lines[pos][x] < lines[y][x]; {
				pos++
				visible["down"]++
			}

			for pos := x - 1; pos > 0 && lines[y][pos] < lines[y][x]; {
				pos--
				visible["left"]++
			}

			for pos := x + 1; pos < len(lines[y])-1 && lines[y][pos] < lines[y][x]; {
				pos++
				visible["right"]++
			}

			score := visible["up"] * visible["down"] * visible["left"] * visible["right"]
			if score > result {
				result = score
			}
		}
	}

	return result
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
