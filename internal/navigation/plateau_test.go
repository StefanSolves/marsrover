package navigation

import "testing"

func TestIsValidCoordinate(t *testing.T) {
	// Initialize a 5x5 grid
	p := NewPlateau(5, 5)

	tests := []struct {
		name string
		x, y int
		want bool
	}{
		{"Bottom Left (Valid)", 0, 0, true},
		{"Top Right (Valid)", 5, 5, true},
		{"Middle (Valid)", 3, 3, true},
		{"X Out of Bounds", 6, 5, false},
		{"Y Out of Bounds", 5, 6, false},
		{"Negative X (Invalid)", -1, 0, false},
		{"Negative Y (Invalid)", 0, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.IsValidCoordinate(tt.x, tt.y); got != tt.want {
				t.Errorf("IsValidCoordinate(%d, %d) = %v, want %v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
