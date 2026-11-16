package main

import "testing"

func TestCalculateDistanceTravelled(t *testing.T) {
	comet := Reindeer{14, 10, 127}
	dancer := Reindeer{16, 11, 162}
	tests := []struct {
		name     string
		reindeer Reindeer
		seconds  int
		want     int
	}{
		{"first", comet, 1, 14},
		{"second", dancer, 1, 16},
		{"third", comet, 10, 140},
		{"fourth", dancer, 10, 160},
		{"fifth", comet, 11, 140},
		{"sixth", dancer, 11, 176},
		{"seventh", comet, 12, 140},
		{"eights", dancer, 12, 176},
		{"nineth", comet, 1000, 1120},
		{"tenth", dancer, 1000, 1056},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateDistanceTravelled(tt.reindeer, tt.seconds)
			if got != tt.want {
				t.Errorf("CalculateDistanceTravelled() = %v, want %v", got, tt.want)
			}
		})
	}
}
