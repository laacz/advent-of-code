package main

import (
	"fmt"
	"os"
	"strings"
)

func lens(s string) (int, int) {
	s = s[1 : len(s)-1]
	alen := len(s)
	elen := alen

	for c := 0; c < len(s); c++ {
		if s[c] == '\\' {
			switch s[c+1] {
			case '\\', '"':
				elen--
				c++
			case 'x':
				elen -= 3
				c += 3
			}
		}
	}

	return alen + 2, elen
}

func encode(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return "\"" + s + "\""
}

func partOne(lines string) int {
	ret := 0
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		alen, elen := lens(line)
		ret += alen - elen
	}
	return ret
}

func partTwo(lines string) int {
	ret := 0
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		line = encode(line)

		alen, elen := lens(line)
		ret += alen - elen
	}
	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
