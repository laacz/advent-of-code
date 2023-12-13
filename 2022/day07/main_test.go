package main

import "testing"

var input string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPartOne(t *testing.T) {
	got := partOne(input)
	expect := 95437
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	got := partTwo(input)
	expect := 24933642
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
