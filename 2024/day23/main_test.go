package main

import (
	"testing"
)

var input string = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

func TestPartOne(t *testing.T) {
	expected := 7
	if actual := part1(parseInput(input)); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := "co,de,ka,ta"
	if actual := part2(parseInput(input)); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
