package main

import "testing"

func Test_parseInstructions(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"first", "(())", 0},
		{"second", "()()", 0},
		{"third", "(((", 3},
		{"fourth", "(()(()(", 3},
		{"fifth", "))(((((", 3},
		{"sixth", "())", -1},
		{"seventh", "))(", -1},
		{"eighth", ")))", -3},
		{"nineth", ")())())", -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInstructions([]byte(tt.in))
			if got != tt.want {
				t.Errorf("parseInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
