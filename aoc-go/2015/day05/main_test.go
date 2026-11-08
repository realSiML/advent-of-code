package main

import "testing"

func TestCheckNiceStringOne(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"first", "ugknbfddgicrmopn", true},
		{"second", "aaa", true},
		{"third", "jchzalrnumimnmhp", false},
		{"fourth", "haegwjzuvuyypxyu", false},
		{"fifth", "dvszwmarrgswjxmb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckNiceStringOne(tt.s)
			if got != tt.want {
				t.Errorf("CheckNiceStringOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckNiceStringTwo(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"first", "qjhvhtzxzqqjkmpb", true},
		{"second", "xxyxx", true},
		{"third", "uurcxstgmygtbstg", false},
		{"fourth", "ieodomkazucvgmuy", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckNiceStringTwo(tt.s)
			if got != tt.want {
				t.Errorf("CheckNiceStringTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
