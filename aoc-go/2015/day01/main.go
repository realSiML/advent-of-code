package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func parseInstructions(s []byte) int {
	var result int

	for _, v := range s {
		switch v {
		case '(':
			result += 1
		case ')':
			result += -1
		}
	}

	return result
}

func PartOne() {
	var floor int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 64*1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			floor += parseInstructions(buf[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Println(floor)
}

func getBasementPosition() func(s []byte) int {
	var result, offset int

	return func(s []byte) int {
		for i, v := range s {
			switch v {
			case '(':
				result += 1
			case ')':
				result += -1
			}

			if result == -1 {
				return i + 1 + offset
			}
		}

		offset += len(s)
		return 0
	}
}

func PartTwo() {
	var position int
	getPos := getBasementPosition()

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 64*1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		position = getPos(buf[:n])
		if position != 0 {
			break
		}
	}

	fmt.Println(position)
}

func main() {
	PartOne()
	PartTwo()
}
