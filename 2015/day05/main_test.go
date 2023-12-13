package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	} {
		actual := isNice(test.input)
		if actual != test.expected {
			t.Errorf("PartOne(%s) = %v, expected %v", test.input, actual, test.expected)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	} {
		actual := isNewNice(test.input)
		if actual != test.expected {
			t.Errorf("PartTwo(%s) = %v, expected %v", test.input, actual, test.expected)
		}
	}
}
