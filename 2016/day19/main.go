package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func parse(input string) (ret int) {
	ret = util.Atoi(util.GetLines(input)[0])

	return
}

func partOne(input string) (ret int) {
	num := parse(input)

	elfs := make([]int, num)

	for i := 0; i < num; i++ {
		elfs[i] = 1
	}

	for {
		total := 0
		for e := range elfs {
			if total > 1 {
				break
			}

			if elfs[e] != 0 {
				ret = e + 1
				total++
			}
		}

		if total == 1 {
			return
		}

		for i := range elfs {
			if elfs[i] == 0 {
				continue
			}
			j := (i + 1) % len(elfs)
			for elfs[j] == 0 && j != i {
				j = (j + 1) % len(elfs)
			}

			elfs[i] += elfs[j]
			elfs[j] = 0
		}
	}
}

func partTwo(input string) (ret int) {
	num := parse(input)

	ret = 1
	for elf := 1; elf < num; elf++ {
		// this is much better (Josephus)
		ret = ret%elf + 1
		if ret > (elf+1)/2 {
			ret++
		}
	}

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
