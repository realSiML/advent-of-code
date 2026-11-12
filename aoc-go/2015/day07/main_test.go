package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		want  string
		want2 Instruction
	}{
		{"first", "123 -> x", "x", Instruction{Kind: SET, Lhs: Operand{"123", true, 123}}},
		{"second", "456 -> y", "y", Instruction{Kind: CONST, Lhs: Operand{"456", true, 456}}},
		{"third", "x AND y -> d", "d", Instruction{Kind: AND, Lhs: Operand{Name: "x"}, Rhs: Operand{Name: "y"}}},
		{"fourth", "x OR y -> e", "e", Instruction{Kind: OR, Lhs: Operand{Name: "x"}, Rhs: Operand{Name: "y"}}},
		{"fifth", "x LSHIFT 2 -> f", "f", Instruction{Kind: LSHIFT, Lhs: Operand{Name: "x"}, Rhs: Operand{"2", true, 2}}},
		{"sixth", "y RSHIFT 2 -> g", "g", Instruction{Kind: RSHIFT, Lhs: Operand{Name: "y"}, Rhs: Operand{"2", true, 2}}},
		{"seventh", "NOT x -> h", "h", Instruction{Kind: NOT, Lhs: Operand{Name: "x"}}},
		{"eights", "NOT y -> i", "i", Instruction{Kind: NOT, Lhs: Operand{Name: "y"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := ParseLine(tt.line)
			if tt.want != got && tt.want2 != got2 {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
