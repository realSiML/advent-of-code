package main

import (
	"fmt"
	"slices"
)

// https://rosettacode.org/wiki/Factors_of_an_integer#Go
func NumberFactors(nr int) []int {
	fs := make([]int, 0)
	if nr < 1 {
		fmt.Println("\nFactors of", nr, "not computed")
		return fs
	}
	fs = append(fs, 1)
	apf := func(p int, e int) {
		n := len(fs)
		for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
			for j := range n {
				fs = append(fs, fs[j]*pp)
			}
		}
	}
	e := 0
	for ; nr&1 == 0; e++ {
		nr >>= 1
	}
	apf(2, e)
	for d := int(3); nr > 1; d += 2 {
		if d*d > nr {
			d = nr
		}
		for e = 0; nr%d == 0; e++ {
			nr /= d
		}
		if e > 0 {
			apf(d, e)
		}
	}
	return fs
}

func PartOne(presents int) int {
	const presentsPerHouse = 10

	house := 1
	for {
		elves := NumberFactors(house)
		elvesSum := 0
		for _, elf := range elves {
			elvesSum += elf
		}
		if elvesSum*presentsPerHouse >= presents {
			break
		}
		house++
	}

	return house
}

func PartTwo(presents int) int {
	const (
		presentsPerHouse = 11
		housesToVisit    = 50
	)

	house := 1
	for {
		elves := NumberFactors(house)
		elves = slices.DeleteFunc(elves, func(elf int) bool {
			return house/elf > housesToVisit
		})
		elvesSum := 0
		for _, elf := range elves {
			elvesSum += elf
		}
		if elvesSum*presentsPerHouse >= presents {
			break
		}
		house++
	}

	return house
}

func main() {
	const input = 34000000

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}
