package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	idx, val int64
}

type File []Pair

// Remove removes a value at a given index.
func (f *File) Remove(i int64) {
	*f = append((*f)[:i], (*f)[i+1:]...)
}

// Insert inserts a value at a given index.
func (f *File) Insert(i int64, p Pair) {
	*f = append((*f)[:i], append(File{p}, (*f)[i:]...)...)
}

// ZeroIndex finds the index of the zero value in the file.
func (f *File) ZeroIndex() int {
	for i, p := range *f {
		if p.val == 0 {
			return i
		}
	}
	return 0
}

// Copy makes a copy of the file.
func (f *File) Copy() File {
	ret := make(File, len(*f))
	copy(ret, *f)
	return ret
}

// FindIndex finds the index of a given value in the file.
func (f *File) FindIndex(p Pair) int64 {
	for i, p1 := range *f {
		if p1 == p {
			return int64(i)
		}
	}
	return 0
}

// Length returns the length of the file.
func (f *File) Length() int64 {
	return int64(len(*f))
}

// Decrypt applies the decryption key to the file.
func (f *File) Decrypt(k int64) {
	for i := range *f {
		(*f)[i].val *= k
	}
}

// Mix does all the actual mixing and returns the result.
func (f *File) Mix(o File) File {
	orig := o.Copy()
	for _, p := range orig {
		if p.val == 0 {
			continue
		}
		idx := f.FindIndex(p)
		newIdx := (idx + p.val) % (orig.Length() - 1)
		if newIdx < 0 {
			newIdx += orig.Length() - 1
		}
		f.Remove(idx)
		f.Insert(newIdx, p)
	}
	return *f
}

// SumCoordinates calculates the result for part1 and part2.
func (f *File) SumCoordinates() int64 {
	idx := f.ZeroIndex()
	return (*f)[(idx+1000)%len(*f)].val + (*f)[(idx+2000)%len(*f)].val + (*f)[(idx+3000)%len(*f)].val
}

// NewFile is just a simple constructor
func NewFile(input string) File {
	var ret File
	for idx, line := range strings.Split(input, "\n") {
		a, _ := strconv.Atoi(line)
		ret = append(ret, Pair{int64(idx), int64(a)})
	}

	return ret
}

func partOne(input string) int64 {
	orig := NewFile(input)
	mix := orig.Mix(orig)
	return mix.SumCoordinates()
}

func partTwo(input string) int64 {
	orig := NewFile(input)
	orig.Decrypt(811589153)
	mix := orig.Copy()
	for i := 0; i < 10; i++ {
		mix = mix.Mix(orig)
	}
	return mix.SumCoordinates()
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
