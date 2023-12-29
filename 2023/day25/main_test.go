package main

import (
	"testing"
)

var tests = []struct {
	input     string
	expected1 int
	expected2 int
}{
	{
		input: `jqt: rhn xhk nvd
		rsh: frs pzl lsr
		xhk: hfx
		cmg: qnr nvd lhk bvb
		rhn: xhk bvb hfx
		bvb: xhk hfx
		pzl: lsr hfx nvd
		qnr: nvd
		ntq: jqt hfx bvb xhk
		nvd: lhk
		lsr: lhk
		rzs: qnr cmg lsr rsh
		frs: qnr lhk lsr`,
		expected1: 54,
		expected2: 47,
	},
}

func TestPartOne(t *testing.T) {
	for _, tt := range tests {
		actual := partOne(tt.input, [][]string{
			{"hfx", "pzl"},
			{"bvb", "cmg"},
			{"nvd", "jqt"},
		})
		if actual != tt.expected1 {
			t.Errorf("Expected %d, got %d", tt.expected1, actual)
		}
	}
}
