package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(line string) (ret []int) {
	for _, line := range strings.Split(line, "\n") {
		var vol int
		fmt.Sscanf(line, "%d", &vol)
		ret = append(ret, vol)
	}

	return ret
}

func combinations(vols []int, liters int, n int) (ret [][]int) {
	if liters == 0 {
		return [][]int{{}}
	}
	if liters < 0 || n == 0 {
		return [][]int{}
	}
	if len(vols) == 0 {
		return [][]int{}
	}

	for i := 0; i < len(vols); i++ {
		vol := vols[i]
		for _, comb := range combinations(vols[i+1:], liters-vol, n-1) {
			ret = append(ret, append([]int{vol}, comb...))
		}
	}

	return ret
}

func partOne(lines string, liters int) (ret int) {
	vols := parse(lines)
	return len(combinations(vols, liters, len(vols)))
}

func partTwo(lines string, liters int) (ret int) {
	ret = 999
	for i := 0; i <= len(parse(lines)); i++ {
		l := len(combinations(parse(lines), liters, i))
		if l != 0 && l < ret {
			ret = l
		}
	}
	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data), 150))
	fmt.Printf("Part two: %d\n", partTwo(string(data), 150))
}
