package main

import (
	"fmt"
	"os"
)

// partOne returns the answer to part one of this day's puzzle
func partOne(lines string) (ret int) {
	row, col := 3010, 3019

	// we need to get position for row, col
	//   | 1 2 3 4
	// --+---------
	// 1 | 1 3 6
	// 2 | 2 5
	// 3 | 4
	// 4 |
	// position is a result of simple arithmetic series where n(i) = n(i-1) + i, starting at one
	r := row + col - 1
	len := r*(r-1)/2 + col

	// now let's loop (3010 * 3019 is small enough for a brute force), starting with 20151125
	ret = 20151125
	for i := 1; i < len; i++ {
		ret = (ret * 252533) % 33554393
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
}
