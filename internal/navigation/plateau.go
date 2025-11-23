package navigation

// Plateau represents the grid on Mars.
// TopRightX and TopRightY are the maximum coordinates (inclusive).
type Plateau struct {
	TopRightX int
	TopRightY int
}

// NewPlateau creates a grid based on the upper-right coordinates.
func NewPlateau(x, y int) *Plateau {
	return &Plateau{
		TopRightX: x,
		TopRightY: y,
	}
}

// IsValidCoordinate checks if a coordinate is within the bounds of the plateau.
// Lower bounds are assumed to be 0,0.
func (p *Plateau) IsValidCoordinate(x, y int) bool {
	return x >= 0 && x <= p.TopRightX && y >= 0 && y <= p.TopRightY
}
