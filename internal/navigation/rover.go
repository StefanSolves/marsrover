package navigation

import "fmt"

// Rover represents the robotic explorer
type Rover struct {
	X       int
	Y       int
	Heading Direction
	Plateau *Plateau
}

// NewRover initialises a rover at a specific position
func NewRover(x, y int, heading Direction, plateau *Plateau) *Rover {
	return &Rover{
		X:       x,
		Y:       y,
		Heading: heading,
		Plateau: plateau,
	}
}

// ExecuteCommands processes a string of characters (L, R, M)
func (r *Rover) ExecuteCommands(commands string) {
	for _, cmd := range commands {
		switch cmd {
		case 'L':
			r.TurnLeft()
		case 'R':
			r.TurnRight()
		case 'M':
			r.Move()
		}
	}
}

func (r *Rover) TurnLeft() {
	r.Heading = r.Heading.Left()
}

func (r *Rover) TurnRight() {
	r.Heading = r.Heading.Right()
}

// Move advances the rover one grid point in the direction it is facing
func (r *Rover) Move() {
	newX, newY := r.X, r.Y

	switch r.Heading {
	case N:
		newY++
	case E:
		newX++
	case S:
		newY--
	case W:
		newX--
	}

	// In a production environment, we check bounds.
	// If the move is invalid, the rover stays put (safety mechanism).
	if r.Plateau.IsValidCoordinate(newX, newY) {
		r.X = newX
		r.Y = newY
	}
}

// CurrentPosition returns the formatted output string (e.g., "1 3 N")
func (r *Rover) CurrentPosition() string {
	return fmt.Sprintf("%d %d %s", r.X, r.Y, r.Heading)
}
