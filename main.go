package main

import (
	"bufio"
	"fmt"
	"marsrover/internal/navigation"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Simple Header
	// We check if it's a terminal just to decide if we print the "Instructions"
	stat, _ := os.Stdin.Stat()
	isInteractive := (stat.Mode() & os.ModeCharDevice) != 0

	if isInteractive {
		fmt.Println("--- Mars Rover Control ---")
		fmt.Println("Paste your input below.")
		fmt.Println("Press ENTER on an empty line to finish.")
		fmt.Println("--------------------------")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var results []string

	// 2. Parse Plateau
	if !scanner.Scan() {
		return
	}
	plateauLine := scanner.Text()
	plateauParams := strings.Fields(plateauLine)
	
	if len(plateauParams) != 2 {
		if isInteractive {
			fmt.Println("Error: Expected Plateau coordinates (e.g. '5 5')")
			return
		}
		panic("Invalid plateau definition")
	}
	
	pX, _ := strconv.Atoi(plateauParams[0])
	pY, _ := strconv.Atoi(plateauParams[1])
	plateau := navigation.NewPlateau(pX, pY)

	// 3. Loop through Rovers
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// THE FIX: Always break on empty line. 
		// This works for manual typing AND files (as long as files don't have random blank lines).
		if trimmedLine == "" {
			break
		}

		parts := strings.Fields(trimmedLine)
		if len(parts) != 3 {
			// Skip malformed lines (like "foo bar")
			continue
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		headingStr := parts[2]

		heading, err := navigation.ParseDirection(headingStr)
		if err != nil {
			continue
		}

		rover := navigation.NewRover(x, y, heading, plateau)

		// Get Commands
		if scanner.Scan() {
			cmdLine := strings.TrimSpace(scanner.Text())
			rover.ExecuteCommands(cmdLine)
			results = append(results, rover.CurrentPosition())
		}
	}

	// 4. Output
	if isInteractive {
		fmt.Println("\n--- Final Rover Positions ---")
	}
	for _, res := range results {
		fmt.Println(res)
	}
}