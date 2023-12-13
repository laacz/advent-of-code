package main

import (
	"fmt"
	"os"
	"strings"
)

func validate(s string) bool {
	// forbidden letters
	if strings.Contains(s, "i") || strings.Contains(s, "o") || strings.Contains(s, "l") {
		return false
	}

	// increasing three letters
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			break
		}
		if i == len(s)-3 {
			return false
		}
	}

	// two pairs
	var pairs int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs++
			i++
		}
		if pairs == 2 {
			break
		}
	}

	return pairs >= 2
}

func next(s string) string {
	buf := []byte(s)
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] == 'z' {
			buf[i] = 'a'
		} else {
			buf[i]++
			break
		}
	}

	return string(buf)
}

func nextValid(s string) string {
	for {
		s = next(s)
		if validate(s) {
			return s
		}
	}
}

func partOne(line string) (ret string) {
	return nextValid(line)
}

func partTwo(line string) (ret string) {
	return nextValid(line)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %s\n", partOne(string(data)))
	fmt.Printf("Part two: %s\n", partTwo(partOne(string(data))))
}
