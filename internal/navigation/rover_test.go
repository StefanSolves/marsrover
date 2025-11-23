package navigation

import "testing"

func TestRoverScenarios(t *testing.T) {
	// Initialize the 5x5 Plateau from the challenge
	plateau := NewPlateau(5, 5)

	tests := []struct {
		name           string
		startX, startY int
		startDir       Direction
		commands       string
		expectedOutput string
	}{
		{
			name:           "Rover 1 (Challenge Data)",
			startX:         1,
			startY:         2,
			startDir:       N,
			commands:       "LMLMLMLMM",
			expectedOutput: "1 3 N",
		},
		{
			name:           "Rover 2 (Challenge Data)",
			startX:         3,
			startY:         3,
			startDir:       E,
			commands:       "MMRMMRMRRM",
			expectedOutput: "5 1 E",
		},
		{
			name:           "Boundary Check (Hit Top Wall)",
			startX:         0,
			startY:         5,
			startDir:       N,
			commands:       "M", // Should not move Y to 6
			expectedOutput: "0 5 N",
		},
		{
			name:     "Boundary Check (Hit Bottom & Left Walls)",
			startX:   0,
			startY:   0,
			startDir: S,
			// 1. Face S, Move (Hit Y=0 wall), Stay 0,0
			// 2. Turn R (Face W), Move (Hit X=0 wall), Stay 0,0
			commands:       "MRM",
			expectedOutput: "0 0 W",
		},
		{
			name:           "360 Degree Turn",
			startX:         1,
			startY:         1,
			startDir:       N,
			commands:       "LLLL", // Full circle Left
			expectedOutput: "1 1 N",
		},
		{
			name:     "Dirty Input (Ignore Unknown Chars)",
			startX:   1,
			startY:   1,
			startDir: N,
			// Should execute M, Ignore X, Execute M. Result: Y+2
			commands:       "MXM",
			expectedOutput: "1 3 N",
		},
		{
			name:           "Empty Command String",
			startX:         2,
			startY:         2,
			startDir:       E,
			commands:       "", // Do nothing
			expectedOutput: "2 2 E",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(tt.startX, tt.startY, tt.startDir, plateau)
			rover.ExecuteCommands(tt.commands)

			if got := rover.CurrentPosition(); got != tt.expectedOutput {
				t.Errorf("Command %s: expected %s, got %s", tt.commands, tt.expectedOutput, got)
			}
		})
	}
}
