package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		presents int
		want     int
	}{
		{"first", 10, 1},
		{"second", 30, 2},
		{"third", 40, 3},
		{"fourth", 70, 4},
		{"fifth", 60, 4},
		{"sixth", 120, 6},
		{"seventh", 80, 6},
		{"eighth", 150, 8},
		{"nineth", 130, 8},
		{"tenth", 20, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartOne(tt.presents)
			if got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
