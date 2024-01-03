package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func parse(input string) (ret [][]int) {
	for _, line := range util.GetLines(input) {
		var a, b, c int
		fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
		ret = append(ret, []int{a, b, c})
	}
	return ret
}

func partOne(input string) (ret int) {
	triangles := parse(input)
	for _, t := range triangles {
		if t[0]+t[1] > t[2] && t[0]+t[2] > t[1] && t[1]+t[2] > t[0] {
			ret++
		}
	}
	return ret
}

func partTwo(input string) (ret int) {
	numbers := parse(input)
	for i := 0; i < len(numbers); i += 3 {
		for j := 0; j < 3; j++ {
			if numbers[i][j]+numbers[i+1][j] > numbers[i+2][j] &&
				numbers[i][j]+numbers[i+2][j] > numbers[i+1][j] &&
				numbers[i+1][j]+numbers[i+2][j] > numbers[i][j] {
				ret++
			}
		}
	}

	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
