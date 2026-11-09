package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want Instruction
	}{
		{"first", "turn on 0,0 through 999,999", Instruction{
			From:   Point{0, 0},
			To:     Point{999, 999},
			Action: TurnOn,
		}},
		{"second", "toggle 0,0 through 999,0", Instruction{
			From:   Point{0, 0},
			To:     Point{999, 0},
			Action: Toggle,
		}},
		{"third", "turn off 499,499 through 500,500", Instruction{
			From:   Point{499, 499},
			To:     Point{500, 500},
			Action: TurnOff,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseLine(tt.line)
			if got != tt.want {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
