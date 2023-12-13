package main

import (
	"testing"
)

func TestValidator(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
		{"abcdefgh", false},
		{"abcdffaa", true},
		{"ghijklmn", false},
	} {
		actual := validate(test.input)
		if actual != test.expected {
			t.Errorf("validate(%s) = %v, expected %v", test.input, actual, test.expected)
		}
	}
}

func TestNextValid(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	} {
		actual := nextValid(test.input)
		if actual != test.expected {
			t.Errorf("next(%s) = %v, expected %v", test.input, actual, test.expected)
		}
	}
}

// func TestPartTwo(t *testing.T) {
// 	expected := 982
// 	actual := partTwo(input)
// 	if actual != expected {
// 		t.Errorf("PartTwo(%s) = %v, expected %v", input, actual, expected)
// 	}
// }
