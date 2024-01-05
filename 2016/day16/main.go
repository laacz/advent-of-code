package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func FillData(a string, length int) (ret string) {
	ret = a
	for len(ret) < length {
		b := make([]byte, len(ret))
		for i := len(ret) - 1; i >= 0; i-- {
			if ret[i] == '0' {
				b[len(b)-i-1] = '1'
			} else {
				b[len(b)-i-1] = '0'
			}
		}
		ret = ret + "0" + string(b)
	}
	return ret[:length]
}

func Checksum(in string) string {
	data := []byte(in)
	for len(data)%2 == 0 {
		newData := []byte{}
		for i := 0; i < len(data); i += 2 {
			if data[i] == data[i+1] {
				newData = append(newData, '1')
			} else {
				newData = append(newData, '0')
			}
		}
		data = newData
	}
	return string(data)
}

func partOne(input string) (ret string) {
	lines := util.GetLines(input)
	data := lines[0]
	length := util.Atoi(lines[1])
	data = FillData(data, length)
	ret = Checksum(data)

	return
}

func partTwo(input string) (ret string) {
	lines := util.GetLines(input)
	data := lines[0]
	length := 35651584
	data = FillData(data, length)
	ret = Checksum(data)

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
