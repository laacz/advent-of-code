package util

import (
	"os"
	"strconv"
	"strings"
)

// ReadFile reads a file and returns its contents as a string ignoring any errors
func ReadFile(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

// GetLines splits a string into lines and returns them as a slice of strings
func GetLines(input string) (ret []string) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ret = append(ret, line)
	}

	return ret
}

// Atoi converts a string to an int ignoring any errors
func Atoi(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

// Abs returns the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
