package main

import (
	"fmt"
	"os"
)

func lookAndSay(s string) (ret string) {
	seq := ""
	for _, c := range s {
		if len(seq) == 0 || seq[0] == byte(c) {
			seq += string(c)
		} else {
			ret += fmt.Sprintf("%d%s", len(seq), string(seq[0]))
			seq = string(c)
		}
	}

	return ret + fmt.Sprintf("%d%s", len(seq), string(seq[0]))
}

func lookAndSay2(s string) (ret string) {
	buf := []byte{}
	seq := ""
	for _, c := range s {
		if len(seq) == 0 || seq[0] == byte(c) {
			seq += string(c)
		} else {
			buf = append(buf, []byte(fmt.Sprintf("%d%s", len(seq), string(seq[0])))...)
			seq = string(c)
		}
	}
	buf = append(buf, []byte(fmt.Sprintf("%d%s", len(seq), string(seq[0])))...)

	return string(buf)
}

func partOne(line string) (ret int) {
	var prevline int = 1
	for i := 0; i < 40; i++ {
		line = lookAndSay2(line)
		fmt.Println(i, len(line), float64(len(line))/float64(prevline))
		prevline = len(line)
	}

	return len(line)
}

func partTwo(line string) (ret int) {
	var prevline int = 1
	for i := 0; i < 50; i++ {
		line = lookAndSay2(line)
		fmt.Println(i, len(line), float64(len(line))/float64(prevline))
		prevline = len(line)
	}

	return len(line)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
