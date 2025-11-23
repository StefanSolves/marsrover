package navigation

import "testing"

func TestParseDirection(t *testing.T) {
	tests := []struct {
		input   string
		want    Direction
		wantErr bool
	}{
		{"N", N, false},
		{"E", E, false},
		{"S", S, false},
		{"W", W, false},
		{"X", 0, true},     // Invalid input
		{"", 0, true},      // Empty input
		{"North", 0, true}, // Wrong format
	}

	for _, tt := range tests {
		got, err := ParseDirection(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseDirection(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			return
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("ParseDirection(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestRotation(t *testing.T) {
	// Ensure the math (iota) wraps around correctly
	tests := []struct {
		start     Direction
		turnLeft  Direction
		turnRight Direction
	}{
		{N, W, E}, // 0 -> 3, 0 -> 1
		{E, N, S}, // 1 -> 0, 1 -> 2
		{S, E, W}, // 2 -> 1, 2 -> 3
		{W, S, N}, // 3 -> 2, 3 -> 0
	}

	for _, tt := range tests {
		if got := tt.start.Left(); got != tt.turnLeft {
			t.Errorf("%s.Left() = %s, want %s", tt.start, got, tt.turnLeft)
		}
		if got := tt.start.Right(); got != tt.turnRight {
			t.Errorf("%s.Right() = %s, want %s", tt.start, got, tt.turnRight)
		}
	}
}
