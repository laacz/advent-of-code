package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func partOne(line string) int {
	ret := 0
	hash := md5.Sum([]byte(line + fmt.Sprintf("%d", ret)))

	for hash[0] != 0 || hash[1] != 0 || hash[2] > 0x0f {
		ret++
		hash = md5.Sum([]byte(line + fmt.Sprintf("%d", ret)))
	}

	return ret
}

func partTwo(line string) int {
	ret := 0
	hash := md5.Sum([]byte(line + fmt.Sprintf("%d", ret)))

	for hash[0] != 0 || hash[1] != 0 || hash[2] > 0x03 {
		ret++
		hash = md5.Sum([]byte(line + fmt.Sprintf("%d", ret)))
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
