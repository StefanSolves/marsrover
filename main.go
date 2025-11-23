package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"marsrover/internal/navigation"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1. Parse Plateau (First Line)
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

	// 2. Loop through Rovers
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
			fmt.Fprintf(os.Stderr, "Error parsing direction: %v\n", err)
			continue
		}

		rover := navigation.NewRover(x, y, heading, plateau)

		if scanner.Scan() {
			cmdLine := strings.TrimSpace(scanner.Text())
			rover.ExecuteCommands(cmdLine)
			fmt.Println(rover.CurrentPosition())
		}
	}
}
