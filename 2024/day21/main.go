package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

type Coord [2]int

type Pad map[rune]Coord

var NumPad = Pad{
	// 	+---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+
	'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
	'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
	'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
	' ': {0, 3}, '0': {1, 3}, 'A': {2, 3},
}

var DirPad = Pad{
	//     +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	' ': {0, 0}, '^': {1, 0}, 'A': {2, 0},
	'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
}

var memo = map[string]int{}

type State struct {
	pos  Coord
	path string
}

func getMoves(pad Pad, start, end rune) []string {
	if start == end {
		return []string{"A"}
	}

	startPos := pad[start]
	endPos := pad[end]
	queue := []State{{pos: startPos, path: ""}}
	distances := make(map[Coord]int)
	var paths []string

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == endPos {
			paths = append(paths, current.path+"A")
			continue
		}

		if dist, exists := distances[current.pos]; exists && dist < len(current.path) {
			continue
		}

		for c, dir := range map[rune]Coord{
			'^': {0, -1},
			'>': {1, 0},
			'v': {0, 1},
			'<': {-1, 0},
		} {
			nextPos := Coord{
				current.pos[0] + dir[0],
				current.pos[1] + dir[1],
			}

			if nextPos == pad[' '] {
				continue
			}

			isButton := false
			for _, buttonPos := range pad {
				if nextPos == buttonPos {
					isButton = true
					break
				}
			}

			if isButton {
				newPath := current.path + string(c)
				if dist, exists := distances[nextPos]; !exists || dist >= len(newPath) {
					queue = append(queue, State{pos: nextPos, path: newPath})
					distances[nextPos] = len(newPath)
				}
			}
		}
	}

	slices.SortFunc(paths, func(a, b string) int {
		return len(a) - len(b)
	})

	return paths
}

func shortestSequence(pad Pad, code string, level int) int {
	key := fmt.Sprintf("%s-%d", code, level)
	if val, ok := memo[key]; ok {
		return val
	}

	current := 'A'
	ret := 0
	for _, c := range code {
		moves := getMoves(pad, current, c)
		if level == 0 {
			ret += len(moves[0])
		} else {
			minLen := math.MaxInt
			for _, move := range moves {
				len := shortestSequence(DirPad, move, level-1)
				if len < minLen {
					minLen = len
				}
			}
			ret += minLen
		}
		current = c
	}
	memo[key] = ret
	return ret
}

func part1(codes []string) int {
	var ret int

	for _, code := range codes {
		length := shortestSequence(NumPad, code, 2)
		var num int
		fmt.Sscanf(code, "%d", &num)
		ret += length * num
	}

	return ret
}

func part2(codes []string) int {
	var ret int

	for _, code := range codes {
		length := shortestSequence(NumPad, code, 25)
		var num int
		fmt.Sscanf(code, "%d", &num)
		ret += length * num
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	maze := parseInput(string(input))
	fmt.Println("part1", part1(maze))
	fmt.Println("part2", part2(maze))
}
