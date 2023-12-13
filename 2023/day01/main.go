package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(lines []string) int {
	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		l := strings.Map(func(r rune) rune {
			if r >= '0' && r <= '9' {
				return r
			}
			return -1
		}, line)

		sum += int(l[0]-'0')*10 + int(l[len(l)-1]-'0')
	}

	return sum
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseDigits(s string) []int {
	var ret []int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ret = append(ret, int(s[i]-'0'))
		}

		for k, v := range digits {
			if strings.HasPrefix(s[i:], k) {
				ret = append(ret, v)
			}
		}
	}

	return ret
}

func partTwo(lines []string) int {
	sum := 0

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		digits := parseDigits(line)

		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum

}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	fmt.Printf("Part one: %d\n", partOne(lines))
	fmt.Printf("Part two: %d\n", partTwo(lines))
}
