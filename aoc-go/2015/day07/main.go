package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Kind int

const (
	CONST Kind = iota
	SET
	AND
	OR
	LSHIFT
	RSHIFT
	NOT
)

type Operand struct {
	Name   string
	IsNum  bool
	Number uint16
}

func ParseOperand(s string) Operand {
	if num, err := strconv.Atoi(s); err == nil {
		return Operand{s, true, uint16(num)}
	}
	return Operand{s, false, 0}
}

type Instruction struct {
	Kind     Kind
	Lhs, Rhs Operand
}

func ParseLine(line string) (target string, instruction Instruction) {
	parts := strings.Split(line, " -> ")
	target = parts[1]

	fields := strings.Fields(parts[0])
	switch len(fields) {
	case 1:
		operand := ParseOperand(fields[0])
		if operand.IsNum {
			instruction.Kind = SET
		} else {
			instruction.Kind = CONST
		}
		instruction.Lhs = operand
	case 2:
		instruction.Kind = NOT
		instruction.Lhs = ParseOperand(fields[1])
	case 3:
		instruction.Lhs = ParseOperand(fields[0])
		instruction.Rhs = ParseOperand(fields[2])
		switch fields[1] {
		case "AND":
			instruction.Kind = AND
		case "OR":
			instruction.Kind = OR
		case "LSHIFT":
			instruction.Kind = LSHIFT
		case "RSHIFT":
			instruction.Kind = RSHIFT
		}
	}

	return target, instruction
}

func Eval(name string, instructions map[string]Instruction, memo map[string]uint16) uint16 {
	if v, err := strconv.Atoi(name); err == nil {
		return uint16(v)
	}

	if v, ok := memo[name]; ok {
		return v
	}

	var v uint16

	inst := instructions[name]
	switch inst.Kind {
	case SET:
		v = Eval(inst.Lhs.Name, instructions, memo)
	case CONST:
		v = Eval(inst.Lhs.Name, instructions, memo)
	case AND:
		v = Eval(inst.Lhs.Name, instructions, memo) & Eval(inst.Rhs.Name, instructions, memo)
	case OR:
		v = Eval(inst.Lhs.Name, instructions, memo) | Eval(inst.Rhs.Name, instructions, memo)
	case LSHIFT:
		v = Eval(inst.Lhs.Name, instructions, memo) << Eval(inst.Rhs.Name, instructions, memo)
	case RSHIFT:
		v = Eval(inst.Lhs.Name, instructions, memo) >> Eval(inst.Rhs.Name, instructions, memo)
	case NOT:
		v = ^Eval(inst.Lhs.Name, instructions, memo)
	}

	memo[name] = v
	return v
}

func main() {
	instructions := map[string]Instruction{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		target, instruction := ParseLine(scanner.Text())
		instructions[target] = instruction
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ans1 := Eval("a", instructions, map[string]uint16{})
	fmt.Println("Part 1:", ans1)

	ans2 := Eval("a", instructions, map[string]uint16{"b": ans1})
	fmt.Println("Part 2:", ans2)
}
