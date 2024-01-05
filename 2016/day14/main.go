package main

import (
	"crypto/md5"
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func partOne(input string) (ret int) {
	keys := []string{}
	hashes := map[int]string{}
	i := 0
	var hash, hash2 string
	var ok bool
	for {
		if hash, ok = hashes[i]; !ok {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
			hashes[i] = hash
		}

		// get slice of three in a row characters
		chars := map[byte]int{}
		for j := 0; j < len(hash)-2; j++ {
			if hash[j] == hash[j+1] && hash[j] == hash[j+2] {
				chars[hash[j]] = 0
				break
			}
		}

		if len(chars) == 0 {
			i++
			continue
		}

	outer:
		for j := i + 1; j < i+1002; j++ {
			if hash2, ok = hashes[j]; !ok {
				hash2 = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, j))))
				hashes[j] = hash2
			}
			for char := range chars {
				for k := 0; k < len(hash2)-4; k++ {
					if hash2[k] == char &&
						hash2[k] == hash2[k+1] && hash2[k] == hash2[k+2] &&
						hash2[k] == hash2[k+3] && hash2[k] == hash2[k+4] {
						keys = append(keys, hash)
						fmt.Print("Found key ", len(keys), ", char ", "\033[1;34m", string(rune(char)), "\033[0m", " at index ", i)
						fmt.Print(", next +", j-i, " hash was ")
						fmt.Print(hash2[:k], "\033[1;34m", hash2[k:k+5], "\033[0m", hash2[k+5:])
						fmt.Println()
						break outer
					}
				}
			}
		}

		if len(keys) == 64 {
			return i
		}
		i++
	}
}

func partTwo(input string) (ret int) {
	keys := []string{}
	hashes := map[int]string{}
	i := 0
	var hash, hash2 string
	var ok bool
	for {
		if hash, ok = hashes[i]; !ok {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
			for q := 0; q < 2016; q++ {
				hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
			}
			hashes[i] = hash
		}

		// get slice of three in a row characters
		chars := map[byte]int{}
		for j := 0; j < len(hash)-2; j++ {
			if hash[j] == hash[j+1] && hash[j] == hash[j+2] {
				chars[hash[j]] = 0
				break
			}
		}

		if len(chars) == 0 {
			i++
			continue
		}

	outer:
		for j := i + 1; j < i+1002; j++ {
			if hash2, ok = hashes[j]; !ok {
				hash2 = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, j))))
				for q := 0; q < 2016; q++ {
					hash2 = fmt.Sprintf("%x", md5.Sum([]byte(hash2)))
				}
				hashes[j] = hash2
			}
			for char := range chars {
				for k := 0; k < len(hash2)-4; k++ {
					if hash2[k] == char &&
						hash2[k] == hash2[k+1] && hash2[k] == hash2[k+2] &&
						hash2[k] == hash2[k+3] && hash2[k] == hash2[k+4] {
						keys = append(keys, hash)
						fmt.Print("Found key ", len(keys), ", char ", "\033[1;34m", string(rune(char)), "\033[0m", " at index ", i)
						fmt.Print(", next +", j-i, " hash was ")
						fmt.Print(hash2[:k], "\033[1;34m", hash2[k:k+5], "\033[0m", hash2[k+5:])
						fmt.Println()
						break outer
					}
				}
			}
		}

		if len(keys) == 64 {
			return i
		}
		i++
	}
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
