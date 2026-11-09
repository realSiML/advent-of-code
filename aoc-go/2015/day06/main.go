package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ROWS = 1000
	COLS = 1000
)

type Action string

const (
	TurnOn  Action = "turn on"
	TurnOff Action = "turn off"
	Toggle  Action = "toggle"
)

func (a Action) IsValid() bool {
	switch a {
	case TurnOn, TurnOff, Toggle:
		return true
	}
	return false
}

type Point struct {
	X int
	Y int
}

type Instruction struct {
	From   Point
	To     Point
	Action Action
}

func ParseLine(line string) Instruction {
	instruction := Instruction{}

	// turn on/turn off/toggle x1,y1 through x2,y2
	fields := strings.Fields(line)

	var i int
	if action := Action(fields[0]); action.IsValid() {
		instruction.Action = Action(fields[0])
		i = 1
	} else {
		action := Action(fields[0]) + " " + Action(fields[1])
		if !action.IsValid() {
			log.Fatalf("Invalid Action: %v", action)
		}

		instruction.Action = action
		i = 2
	}

	from := strings.Split(fields[i], ",")
	x, err := strconv.Atoi(from[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(from[1])
	if err != nil {
		log.Fatal(err)
	}
	instruction.From = Point{x, y}

	i += 2

	to := strings.Split(fields[i], ",")
	x, err = strconv.Atoi(to[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err = strconv.Atoi(to[1])
	if err != nil {
		log.Fatal(err)
	}
	instruction.To = Point{x, y}

	return instruction
}

func ApplyInstruction(grid *[ROWS][COLS]bool, instruction Instruction) {
	x1, y1 := instruction.From.X, instruction.From.Y
	x2, y2 := instruction.To.X, instruction.To.Y
	action := instruction.Action

	for i := y1; i <= y2; i++ {
		for j := x1; j <= x2; j++ {
			switch action {
			case Toggle:
				grid[i][j] = !grid[i][j]
			case TurnOn:
				grid[i][j] = true
			case TurnOff:
				grid[i][j] = false
			}
		}
	}
}

func CountLights(grid [ROWS][COLS]bool) int {
	var count int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				count++
			}
		}
	}

	return count
}

func PartOne() {
	var grid [ROWS][COLS]bool
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			instruction := ParseLine(line)
			ApplyInstruction(&grid, instruction)
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(CountLights(grid))
}

func ApplyInstructionNew(grid *[ROWS][COLS]int, instruction Instruction) {
	x1, y1 := instruction.From.X, instruction.From.Y
	x2, y2 := instruction.To.X, instruction.To.Y
	action := instruction.Action

	for i := y1; i <= y2; i++ {
		for j := x1; j <= x2; j++ {
			switch action {
			case Toggle:
				grid[i][j] += 2
			case TurnOn:
				grid[i][j]++
			case TurnOff:
				if grid[i][j] > 0 {
					grid[i][j]--
				}
			}
		}
	}
}

func CountBrightness(grid [ROWS][COLS]int) int {
	var brightness int

	for i := range grid {
		for j := range grid[i] {
			brightness += grid[i][j]
		}
	}

	return brightness
}

func PartTwo() {
	var grid [ROWS][COLS]int
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			instruction := ParseLine(line)
			ApplyInstructionNew(&grid, instruction)
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	fmt.Println(CountBrightness(grid))
}

func main() {
	PartOne()
	PartTwo()
}
