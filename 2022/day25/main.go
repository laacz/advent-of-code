package main

import (
	"fmt"
	"os"
	"strings"
)

func snafuToDecimal(snafu string) int64 {
	var snafuMap = map[uint8]int64{
		'=': -2,
		'-': -1,
		'0': 0,
		'1': 1,
		'2': 2,
	}
	var ret, factor int64 = 0, 1
	for i := len(snafu) - 1; i >= 0; i-- {
		ret += snafuMap[snafu[i]] * factor
		factor *= 5
	}
	return ret
}

func decimalToSnafu(decimal int64) string {
	var snafuMap = map[int64]string{
		-2: "=",
		-1: "-",
		0:  "0",
		1:  "1",
		2:  "2",
	}
	ret, dec := "", decimal
	for dec != 0 {
		rem := dec % 5
		dec = dec / 5
		if rem >= 3 {
			dec, rem = dec+1, rem-5
		}
		ret = snafuMap[rem] + ret
	}

	return ret
}

func partOne(input string) string {
	var ret int64
	for _, line := range strings.Split(input, "\n") {
		ret += snafuToDecimal(line)
	}
	return decimalToSnafu(ret)
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Final answer: %s\n", partOne(string(data)))
}
