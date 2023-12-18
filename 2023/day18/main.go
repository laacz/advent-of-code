package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Area returns the area of the polygon using Shoelace formula, adding area of the perimeter line
func area(vertices [][]int) (ret int) {
	// this takes into account that vertices are given in an ordered fashion
	var i int
	for i = 0; i < int(len(vertices)-1); i++ {
		ret += vertices[i][1]*vertices[i+1][0] - vertices[i][0]*vertices[i+1][1]
	}

	ret += vertices[len(vertices)-1][1]*vertices[0][0] - vertices[len(vertices)-1][0]*vertices[0][1]
	for i := 0; i < len(vertices); i++ {
		ret += vertices[i][2]
	}

	ret = ret/2 + 1

	return ret
}

// parse parses the input into a Grid according to the rules of the first part of the puzzle
func parse(input string) (ret [][]int) {
	var x, y int
	for _, line := range strings.Split(input, "\n") {
		var dir rune
		var steps int
		var color string
		fmt.Sscanf(line, "%c %d (#%s)", &dir, &steps, &color)

		switch dir {
		case 'R':
			x += steps
		case 'D':
			y += steps
		case 'L':
			x -= steps
		case 'U':
			y -= steps
		}
		ret = append(ret, []int{y, x, steps})
	}

	return ret
}

// parse2 parses the input into a Grid according to the rules of the second part of the puzzle
func parse2(input string) (ret [][]int) {
	var x, y int
	for _, line := range strings.Split(input, "\n") {
		var dir rune
		var steps int
		var color string
		fmt.Sscanf(line, "%c %d (#%s)", &dir, &steps, &color)

		tmp, _ := strconv.ParseInt(color[:5], 16, 32)
		steps = int(tmp)
		dir = rune(color[5])

		switch dir {
		case '0':
			x += steps
		case '1':
			y += steps
		case '2':
			x -= steps
		case '3':
			y -= steps
		}
		ret = append(ret, []int{y, x, steps})
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	return area(parse(input))
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	return area(parse2(input))
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
