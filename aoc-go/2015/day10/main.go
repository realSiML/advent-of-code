package main

import (
	"fmt"
	"strconv"
	"strings"
)

func LookAndSay(s string) string {
	if s == "" {
		return ""
	}

	var sb strings.Builder
	count := 1

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
			continue
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(s[i-1])
		count = 1
	}

	return sb.String()
}

func main() {
	const INPUT = "1321131112"
	input := INPUT

	for range 40 {
		input = LookAndSay(input)
	}
	partOne := len(input)
	fmt.Println(partOne)

	for range 10 {
		input = LookAndSay(input)
	}
	partTwo := len(input)
	fmt.Println(partTwo)
}
