package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

type Room struct {
	name, checksum string
	sectorID       int
	actualCecksum  string
}

func (r Room) IsValid() bool {
	return r.checksum == r.actualCecksum
}

func (r Room) Decrypt() string {
	ret := ""
	for _, c := range r.name {
		if c == '-' {
			ret += " "
		} else {
			ret += string(byte('a' + (int(c-'a')+r.sectorID)%26))
		}
	}
	return ret
}

type LetterRank struct {
	letter rune
	count  int
}

func (lr LetterRank) String() string {
	return fmt.Sprintf("%c: %d", lr.letter, lr.count)
}

func NewRoom(line string) (ret Room) {
	parts := strings.Split(line, "[")
	ret.checksum = parts[1][:len(parts[1])-1]

	parts = strings.Split(parts[0], "-")
	ret.name = strings.Join(parts[:len(parts)-1], "-")
	ret.sectorID = util.Atoi(parts[len(parts)-1])

	rank := []LetterRank{}
	for _, c := range parts[:len(parts)-1] {
		for _, r := range c {
			added := false
			for ii, rr := range rank {
				if rr.letter == r {
					rank[ii].count++
					added = true
				}
			}
			if !added {
				rank = append(rank, LetterRank{r, 1})
			}
		}
	}

	sort.Slice(rank, func(i, j int) bool {
		if rank[i].count == rank[j].count {
			return rank[i].letter < rank[j].letter
		}
		return rank[i].count > rank[j].count
	})

	for _, r := range rank[:5] {
		ret.actualCecksum += string(r.letter)
	}

	return ret
}

func parse(input string) (ret []Room) {
	for _, line := range util.GetLines(input) {
		ret = append(ret, NewRoom(line))
	}
	return ret
}

func partOne(input string) (ret int) {
	rooms := parse(input)
	for _, room := range rooms {
		if room.IsValid() {
			ret += room.sectorID
		}
	}
	return ret
}

func partTwo(input string) (ret int) {
	rooms := parse(input)
	for _, room := range rooms {
		if room.Decrypt() == "northpole object storage" {
			ret = room.sectorID
			break
		}
	}

	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
