package main

import "testing"

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name string
		p    string
		want bool
	}{
		{"first", "hijklmmn", false},
		{"second", "abbceffg", false},
		{"third", "abbcegjk", false},
		{"fourth", "abcdffaa", true},
		{"five", "ghjaabcc", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.p)
			if got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextPassword(t *testing.T) {
	tests := []struct {
		name string
		p    string
		want string
	}{
		{"first", "abcdefg", "abcdefh"},
		{"second", "abcdefz", "abcdega"},
		{"third", "azzzzzz", "baaaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNextPassword(tt.p)
			if got != tt.want {
				t.Errorf("GetNextPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
