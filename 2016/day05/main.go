package main

import (
	"crypto/md5"
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func GuessPassword(input string) (ret string) {
	i := 0
	for len(ret) < 8 {
		h := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		if h[0] == 0 && h[1] == 0 && h[2] < 16 {
			d := fmt.Sprintf("%x", h[2])
			ret += d
			fmt.Println("Next character:", d)
		}
		i++
	}

	return ret
}

func GuessPasswordWityhPos(input string) (ret string) {
	i := 0
	r := []byte{255, 255, 255, 255, 255, 255, 255, 255}
outer:
	for {
		h := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))

		if h[0] == 0 && h[1] == 0 && h[2] < 16 {
			pos := h[2]
			if pos < 8 && r[pos] == 255 {
				r[pos] = h[3] >> 4
				fmt.Printf("Found character at position %d: %x\n", pos, r[pos])
			}
		}
		i++

		for _, v := range r {
			if v == 255 {
				continue outer
			}
		}
		break
	}

	for _, v := range r {
		ret += fmt.Sprintf("%x", v)
	}

	return ret
}

func partOne(input string) (ret string) {
	ret = GuessPassword(input)
	return ret
}

func partTwo(input string) (ret string) {
	ret = GuessPasswordWityhPos(input)
	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
