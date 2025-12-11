package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := part1(parse(string(data)))
	fmt.Println("Part 1:", result)

	result2 := part2(parse(string(data)))
	fmt.Println("Part 2:", result2)
}

type Instruction struct {
	address int
	value   int
	mask    string
}

func parse(input string) []Instruction {
	ret := []Instruction{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		var instr Instruction
		if strings.HasPrefix(line, "mask") {
			instr = Instruction{
				mask: line[7:],
			}
		} else if strings.HasPrefix(line, "mem") {
			parts := strings.Split(line, " = ")
			addrPart := parts[0]
			valuePart := parts[1]

			addrStr := addrPart[strings.Index(addrPart, "[")+1 : strings.Index(addrPart, "]")]
			address, _ := strconv.Atoi(addrStr)
			value, _ := strconv.Atoi(valuePart)

			instr = Instruction{
				address: address,
				value:   value,
			}
		} else {
			continue
		}

		ret = append(ret, instr)
	}

	return ret
}

func part1(input []Instruction) int {
	mask := ""
	memory := make(map[int]int)
	for _, instr := range input {
		if instr.mask != "" {
			mask = instr.mask
		} else {
			value := instr.value
			valueBin := fmt.Sprintf("%036b", value)
			maskedValue := ""
			for i := 0; i < 36; i++ {
				if mask[i] == 'X' {
					maskedValue += string(valueBin[i])
				} else {
					maskedValue += string(mask[i])
				}
			}
			finalValue, _ := strconv.ParseInt(maskedValue, 2, 64)
			memory[instr.address] = int(finalValue)
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}

	return sum
}

func expand(masked string) []int {
	idx := strings.Index(masked, "X")
	if idx == -1 {
		addr, _ := strconv.ParseInt(masked, 2, 64)
		return []int{int(addr)}
	}

	with0 := masked[:idx] + "0" + masked[idx+1:]
	with1 := masked[:idx] + "1" + masked[idx+1:]

	return append(expand(with0), expand(with1)...)
}

func part2(input []Instruction) int {
	mask := ""
	memory := make(map[int]int)

	for _, instr := range input {
		if instr.mask != "" {
			mask = instr.mask
		} else {
			addrBin := fmt.Sprintf("%036b", instr.address)
			maskedAddr := ""
			for i := 0; i < 36; i++ {
				switch mask[i] {
				case '0':
					maskedAddr += string(addrBin[i])
				case '1':
					maskedAddr += "1"
				case 'X':
					maskedAddr += "X"
				}
			}

			addresses := expand(maskedAddr)
			for _, addr := range addresses {
				memory[addr] = instr.value
			}
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}
