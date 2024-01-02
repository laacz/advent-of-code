package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expect := 2
	p := parse(`inc a
	jio a, +2
	tpl a
	inc a`)
	got := p.Run(Registers{"a": 0, "b": 0})["a"]
	if got != expect {
		t.Errorf("PartOne() = %v, want %v", got, expect)
	}
}
