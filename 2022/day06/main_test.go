package main

import "testing"

func TestPartOne(t *testing.T) {
	var tests = map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for input, expect := range tests {
		got := partOne(input)
		if got != expect {
			t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
		}
	}
}

func TestPartTwo(t *testing.T) {
	var tests = map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}

	for input, expect := range tests {
		got := partTwo(input)
		if got != expect {
			t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
		}
	}
}
