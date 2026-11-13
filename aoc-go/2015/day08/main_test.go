package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		wantTotal    int
		wantInMemory int
	}{
		{"first", "\"\"", 2, 0},
		{"second", "\"abc\"", 5, 3},
		{"third", "\"aaa\\\"aaa\"", 10, 7},
		{"fourth", "\"\\x27\"", 6, 1},
		{"fifth", "\"dmlpcglcfkfsctxydjvayhymv\\x3c\\\\gp\"", 35, 29},
		{"sixth", "\"nrhddblbuzlqsl\\x9cben\"", 23, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotInMemory := ParseLine(tt.line)
			if gotTotal != tt.wantTotal || gotInMemory != tt.wantInMemory {
				t.Errorf("ParseLine() = %v, %v; want %v, %v", gotTotal, gotInMemory, tt.wantTotal, tt.wantInMemory)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"first", `""`, `"\"\""`},
		{"second", `"abc"`, `"\"abc\""`},
		{"third", `"aaa\"aaa"`, `"\"aaa\\\"aaa\""`},
		{"fourth", `"\x27"`, `"\"\\x27\""`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.s)
			if got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
