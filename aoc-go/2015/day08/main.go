package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseLine(line string) (int, int) {
	var total, inMemory int
	var escaped, hex bool
	var hexCount int

	for i := range line {
		total++

		if i == 0 || i == len(line)-1 {
			continue
		}

		c := line[i]

		if !escaped {
			if c == '\\' {
				escaped = true
				continue
			}
		}

		if escaped {
			if hex {
				hexCount++
				if hexCount == 2 {
					inMemory++
					hexCount = 0
					hex = false
					escaped = false
				}
				continue
			}

			if c == 'x' {
				hex = true
				continue
			}

			if c == '\\' || c == '"' {
				inMemory++
				escaped = false
				continue
			}
		}

		inMemory++
	}

	return total, inMemory
}

func Encode(s string) string {
	var builder strings.Builder

	builder.WriteByte('"')
	for i := range s {
		c := s[i]
		switch c {
		case '"':
			builder.WriteString("\\\"")
		case '\\':
			builder.WriteString("\\\\")
		default:
			builder.WriteByte(c)
		}
	}
	builder.WriteByte('"')

	return builder.String()
}

func main() {
	var partOne, partTwo int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		v1, v2 := ParseLine(line)
		partOne += v1 - v2

		encodedLine := Encode(line)
		v1, v2 = ParseLine(encodedLine)
		partTwo += v1 - v2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne)
	fmt.Println(partTwo)
}
