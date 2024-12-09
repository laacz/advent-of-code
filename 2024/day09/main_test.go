package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect []int
	}{
		{"12345", toSlice("0..111....22222")},
		{"2333133121414131402", toSlice("00...111...2...333.44.5555.6666.777.888899")},
	} {
		actual := parseInput(c.input)
		if !reflect.DeepEqual(actual, c.expect) {
			t.Errorf("Expected %+v, got %+v", c.expect, actual)
		}
	}
}

func TestDefragment(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect []int
	}{
		{toSlice("0..111....22222"), toSlice("022111222......")},
		{toSlice("00...111...2...333.44.5555.6666.777.888899"), toSlice("0099811188827773336446555566..............")},
		{toSlice("00...1111111999999999999"), toSlice("009991111111999999999...")},
	} {
		actual := DefragmentBlocks(c.input)
		if !reflect.DeepEqual(actual, c.expect) {
			t.Errorf("Expected %+v, got %+v", c.expect, actual)
		}
	}
}

func TestPartOne(t *testing.T) {
	input := "2333133121414131402"
	expect := 1928
	actual := part1(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := "2333133121414131402"
	expect := 2858
	actual := part2(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func toSlice(s string) []int {
	var ret []int
	for _, c := range s {
		if c == '.' {
			ret = append(ret, -1)
			continue
		}
		ret = append(ret, int(c-'0'))
	}
	return ret
}
