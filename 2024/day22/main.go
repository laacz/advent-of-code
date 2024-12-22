package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	var ret []int

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		n, _ := strconv.Atoi(line)
		ret = append(ret, n)
	}

	return ret
}

func prune(sn int) int {
	return sn % 16777216
}

func mix(sn, val int) int {
	return sn ^ val
}

func nextSn(sn int) int {
	ret := prune(mix(sn, sn*64))
	ret = prune(mix(ret, ret/32))
	ret = prune(mix(ret, ret*2048))

	return ret
}

func part1(sns []int) int {
	var ret int

	for _, sn := range sns {
		newSn := sn
		for i := 0; i < 2000; i++ {
			newSn = nextSn(newSn)
		}
		ret += newSn
	}

	return ret
}

func part2(sns []int) int {
	var ret int

	sequences := make(map[[4]int]int)

	for _, sn := range sns {
		prevSn, newSn := sn, sn

		seq := [4]int{}

		seenSeqs := make(map[[4]int]bool)
		for i := 0; i < 2000; i++ {
			newSn = nextSn(newSn)
			price := newSn % 10
			seq = [4]int{
				seq[1],
				seq[2],
				seq[3],
				newSn%10 - prevSn%10,
			}
			prevSn = newSn
			if seenSeqs[seq] || i < 3 {
				continue
			}
			seenSeqs[seq] = true
			sequences[seq] += price
		}
	}

	for _, v := range sequences {
		if v > ret {
			ret = v
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
