package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func parse(input string) (ret []string) {
	ret = append(ret, util.GetLines(input)...)
	return ret
}

var dirs = map[rune][2]int{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
}

func partOne(input string) (ret string) {
	code := parse(input)

	pad := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	pos := [2]int{1, 1}
	for _, v := range code {
		for _, c := range v {
			np := [2]int{pos[0] + dirs[c][0], pos[1] + dirs[c][1]}
			if np[0] < 0 || np[0] > 2 || np[1] < 0 || np[1] > 2 {
				continue
			}
			pos = np
		}
		ret += fmt.Sprintf("%d", pad[pos[1]][pos[0]])
	}

	return ret
}

func partTwo(input string) (ret string) {
	code := parse(input)

	pad := [5][5]byte{
		{'0', '0', '1', '0', '0'},
		{'0', '2', '3', '4', '0'},
		{'5', '6', '7', '8', '9'},
		{'0', 'A', 'B', 'C', '0'},
		{'0', '0', 'D', '0', '0'},
	}

	pos := [2]int{0, 2}
	for _, v := range code {
		for _, c := range v {
			np := [2]int{pos[0] + dirs[c][0], pos[1] + dirs[c][1]}
			if np[0] < 0 || np[0] > 4 || np[1] < 0 || np[1] > 4 || pad[np[1]][np[0]] == '0' {
				continue
			}
			pos = np
		}
		ret += string(rune(pad[pos[1]][pos[0]]))
	}

	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
