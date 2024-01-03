package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func Decompress(str string, partTwo bool) (ret int) {
	i := 0

	for i < len(str) {
		if str[i] == '(' {
			i += 1

			var marker string
			for str[i] != ')' {
				marker += string(str[i])
				i += 1
			}

			var length, amount int
			fmt.Sscanf(marker, "%dx%d", &length, &amount)

			if partTwo {
				ret += Decompress(str[i+1:i+1+length], true) * amount
			} else {
				ret += amount * length
			}

			i += length
		} else {
			ret += 1
		}
		i += 1
	}

	return ret
}

func partOne(input string) (ret int) {
	ret = Decompress(input, false)
	return ret
}

func partTwo(input string) (ret int) {
	ret = Decompress(input, true)
	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
