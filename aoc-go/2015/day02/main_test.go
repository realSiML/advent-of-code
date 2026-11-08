package main

import "testing"

func TestCalculateRequiredWrapping(t *testing.T) {
	tests := []struct {
		name string
		l    int
		w    int
		h    int
		want int
	}{
		{"first", 2, 3, 4, 58},
		{"second", 1, 1, 10, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateRequiredWrapping(tt.l, tt.w, tt.h)
			if got != tt.want {
				t.Errorf("CalculateRequiredWrapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequiredRibbon(t *testing.T) {
	tests := []struct {
		name string
		l    int
		w    int
		h    int
		want int
	}{
		{"first", 2, 3, 4, 34},
		{"second", 1, 1, 10, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateRequiredRibbon(tt.l, tt.w, tt.h)
			if got != tt.want {
				t.Errorf("CalculateRequiredRibbon() = %v, want %v", got, tt.want)
			}
		})
	}
}
