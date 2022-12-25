package day2

import (
	"advent_of_code/myutils"
	"fmt"
	"strings"
)

func calcPoints(choices []string) int {
	points := 0
	// Based on your choice
	switch choices[1] {
	case "X": // Rock
		points = 1
		switch choices[0] {
		case "A":
			points += 3
		case "B":
			points += 0
		case "C":
			points += 6
		}
	case "Y": // Paper
		points = 2
		switch choices[0] {
		case "A":
			points += 6
		case "B":
			points += 3
		case "C":
			points += 0
		}
	case "Z": // Scissors
		points = 3
		switch choices[0] {
		case "A":
			points += 0
		case "B":
			points += 6
		case "C":
			points += 3
		}
	}
	return points
}

func calcPointsV2(choices []string) int {
	points := 0
	// Based on desired result
	switch choices[1] {
	case "Z": // Win
		points = 6
		switch choices[0] {
		case "A":
			points += 2
		case "B":
			points += 3
		case "C":
			points += 1
		}
	case "Y": // Draw
		points = 3
		switch choices[0] {
		case "A":
			points += 1
		case "B":
			points += 2
		case "C":
			points += 3
		}
	case "X": // Lose
		points = 0
		switch choices[0] {
		case "A":
			points += 3
		case "B":
			points += 1
		case "C":
			points += 2
		}
	}
	return points
}

func Day2() {
	lines := myutils.ReadFileLines("./day2/input")

	totalPoints := 0
	totalPoints2 := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		totalPoints += calcPoints(choices)
		totalPoints2 += calcPointsV2(choices)
	}
	fmt.Printf("Total points: %d\n", totalPoints)
	fmt.Printf("Total points V2: %d\n", totalPoints2)
}
