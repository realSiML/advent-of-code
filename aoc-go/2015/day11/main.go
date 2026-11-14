package main

import (
	"fmt"
	"strings"
)

func ValidatePassword(p string) bool {
	if len(p) < 5 {
		return false
	}

	if strings.ContainsAny(p, "iol") {
		return false
	}

	hasStraightAt := func(p string, i int) bool {
		return i+2 < len(p) &&
			p[i]+1 == p[i+1] &&
			p[i+1]+1 == p[i+2]
	}
	var straight bool
	var pairCount int

	i := 0
	for i < len(p)-1 {
		if hasStraightAt(p, i) {
			straight = true
		}

		if p[i] == p[i+1] {
			pairCount++

			if hasStraightAt(p, i+1) {
				straight = true
			}

			i += 2
		} else {
			i++
		}

	}

	return straight && pairCount >= 2
}

func GetNextPassword(p string) string {
	const start, end byte = 'a', 'z'
	chars := []byte(p)

	chars[len(chars)-1]++
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] <= end {
			break
		}

		chars[i] = chars[i]%end + start - 1
		if i > 0 {
			chars[i-1]++
		}
	}

	return string(chars)
}

func GetNextValidPassword(p string) string {
	p = GetNextPassword(p)
	for !ValidatePassword(p) {
		p = GetNextPassword(p)
	}
	return p
}

func main() {
	const INPUT = "cqjxjnds"

	partOne := GetNextValidPassword(INPUT)
	fmt.Println(partOne)
	partTwo := GetNextValidPassword(partOne)
	fmt.Println(partTwo)
}
