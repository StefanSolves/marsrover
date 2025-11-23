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
	// 1. Check if input is coming from a terminal (Interactive) or a file (Pipe)
	stat, _ := os.Stdin.Stat()
	isInteractive := (stat.Mode() & os.ModeCharDevice) != 0

	if isInteractive {
		fmt.Println("--- Mars Rover Control ---")
		fmt.Println("Enter data (Plateau size, then Rover pairs).")
		fmt.Println("Press ENTER on an empty line to finish and view results.")
		fmt.Println("--------------------------")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var results []string // Buffer to store output

	// 2. Parse Plateau
	if !scanner.Scan() {
		return
	}
	plateauLine := scanner.Text()
	plateauParams := strings.Fields(plateauLine)
	if len(plateauParams) != 2 {
		panic("Invalid plateau definition. Expected two integers (e.g., '5 5')")
	}

	pX, _ := strconv.Atoi(plateauParams[0])
	pY, _ := strconv.Atoi(plateauParams[1])
	plateau := navigation.NewPlateau(pX, pY)

	// 3. Loop through Rovers
	for scanner.Scan() {
		posLine := strings.TrimSpace(scanner.Text())
		if posLine == "" {
			continue
		}

		parts := strings.Fields(posLine)
		if len(parts) != 3 {
			continue
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		headingStr := parts[2]

		heading, err := navigation.ParseDirection(headingStr)
		if err != nil {
			// In interactive mode, we might want to warn the user
			if isInteractive {
				fmt.Printf("Warning: Invalid direction '%s'. Skipping rover.\n", headingStr)
			}
			continue
		}

		rover := navigation.NewRover(x, y, heading, plateau)

		if scanner.Scan() {
			cmdLine := strings.TrimSpace(scanner.Text())
			rover.ExecuteCommands(cmdLine)
			
			// Store result in buffer instead of printing immediately
			results = append(results, rover.CurrentPosition())
		}
	}

	// 4. Print all results at once (Prevents visual clutter)
	if isInteractive {
		fmt.Println("\n--- Final Rover Positions ---")
	}
	for _, res := range results {
		fmt.Println(res)
	}
}