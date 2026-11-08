package main

import "testing"

func TestParseMove(t *testing.T) {
	tests := []struct {
		name  string
		move  string
		wantX int
		wantY int
	}{
		{"left", "<", -1, 0},
		{"right", ">", 1, 0},
		{"up", "^", 0, 1},
		{"down", "v", 0, -1},
		{"other", "\n", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := ParseMove(tt.move)
			if gotX != tt.wantX && gotY != tt.wantY {
				t.Errorf("ParseMove() = %v, %v; want %v, %v", gotX, gotX, tt.wantY, tt.wantX)
			}
		})
	}
}
