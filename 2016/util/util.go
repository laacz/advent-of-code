package util

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

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

func Atoi(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
