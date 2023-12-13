package main

import (
	"strings"
	"testing"
)

var input1 string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var input2 string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

var lines1 = strings.Split(input1, "\n")
var lines2 = strings.Split(input2, "\n")

func TestPartOne(t *testing.T) {
	expect := 142
	actual := partOne(lines1)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 281
	actual := partTwo(lines2)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
