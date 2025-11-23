package navigation

import "fmt"

// Direction represents the compass heading (N, E, S, W)
type Direction int

// Enum constants for directions
const (
	N Direction = iota // 0
	E                  // 1
	S                  // 2
	W                  // 3
)

// String returns the string representation of the direction
func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

// ParseDirection converts a string character to a Direction
func ParseDirection(char string) (Direction, error) {
	switch char {
	case "N":
		return N, nil
	case "E":
		return E, nil
	case "S":
		return S, nil
	case "W":
		return W, nil
	default:
		return 0, fmt.Errorf("invalid direction: %s", char)
	}
}

// Left returns the direction 90 degrees to the left
func (d Direction) Left() Direction {
	// (0 - 1 + 4) % 4 = 3 (W)
	return (d - 1 + 4) % 4
}

// Right returns the direction 90 degrees to the right
func (d Direction) Right() Direction {
	// (0 + 1) % 4 = 1 (E)
	return (d + 1) % 4
}
