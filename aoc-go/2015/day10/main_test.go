package main

import "testing"

func TestLookAndSay(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"first", "1", "11"},
		{"second", "11", "21"},
		{"third", "21", "1211"},
		{"fourth", "1211", "111221"},
		{"fifth", "111221", "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LookAndSay(tt.s)
			if got != tt.want {
				t.Errorf("LookAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
