package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

func ParseMove(move string) (int, int) {
	switch move {
	case "<":
		return -1, 0
	case ">":
		return 1, 0
	case "^":
		return 0, 1
	case "v":
		return 0, -1
	default:
		return 0, 0
	}
}

func PartOne() {
	visited := make(map[Point]bool)
	curPoint := Point{0, 0}
	visited[curPoint] = true

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 64*1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			for _, v := range buf[:n] {
				dx, dy := ParseMove(string(v))
				curPoint.X += dx
				curPoint.Y += dy
				visited[curPoint] = true
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(len(visited))
}

func PartTwo() {
	visited := make(map[Point]bool)
	curPoint1, curPoint2 := Point{0, 0}, Point{0, 0}
	visited[curPoint1] = true

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 64*1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			for i, v := range buf[:n] {
				dx, dy := ParseMove(string(v))
				if i%2 == 0 {
					curPoint1.X += dx
					curPoint1.Y += dy
					visited[curPoint1] = true
				} else {
					curPoint2.X += dx
					curPoint2.Y += dy
					visited[curPoint2] = true
				}
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(len(visited))
}

func main() {
	PartOne()
	PartTwo()
}
