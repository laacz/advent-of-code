package main

import (
	"testing"
)

var INPUT = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func TestSNAFU(t *testing.T) {
	tests := map[string]int64{
		"1=-0-2": 1747,
		"12111":  906,
		"2=0=":   198,
		"21":     11,
		"2=01":   201,
		"111":    31,
		"20012":  1257,
		"112":    32,
		"1=-1=":  353,
		"1-12":   107,
		"12":     7,
		"1=":     3,
		"122":    37,
	}

	for snafu, decimal := range tests {
		if decimal != snafuToDecimal(snafu) {
			t.Errorf("snafuToDecimal(%s) got %d want %d", snafu, snafuToDecimal(snafu), decimal)
		}
		if snafu != decimalToSnafu(decimal) {
			t.Errorf("decimalToSnafu(%d) got %s want %s", decimal, decimalToSnafu(decimal), snafu)
		}
	}
}

func TestPartOne(t *testing.T) {
	got := partOne(INPUT)
	expect := "2=-1=0"
	if got != expect {
		t.Errorf("partOne(...) = %s, want %s", got, expect)
	}
}
