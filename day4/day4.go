package day4

import (
	"advent_of_code/myutils"
	"fmt"
	"strconv"
	"strings"
)

func atoiCheck(in string) int {
	out, err := strconv.Atoi(in)
	myutils.Check(err)
	return out
}

func checkAssignments(assignments []string) int {
	hours1 := strings.Split(assignments[0], "-")
	hours2 := strings.Split(assignments[1], "-")
	s1 := atoiCheck(hours1[0])
	s2 := atoiCheck(hours2[0])
	e1 := atoiCheck(hours1[1])
	e2 := atoiCheck(hours2[1])
	if (s1 >= s2 && e1 <= e2) || (s2 >= s1 && e2 <= e1) {
		return 1
	}
	return 0
}

func checkSlightAssignments(assignments []string) int {
	hours1 := strings.Split(assignments[0], "-")
	hours2 := strings.Split(assignments[1], "-")
	s1 := atoiCheck(hours1[0])
	s2 := atoiCheck(hours2[0])
	e1 := atoiCheck(hours1[1])
	e2 := atoiCheck(hours2[1])
	if (s2 >= s1 && s2 <= e1) || (s1 >= s2 && s1 <= e2) {
		return 1
	}
	return 0
}

func Day4() {
	lines := myutils.ReadFileLines("./day4/input")

	fullOverlaps := 0
	slightOverlaps := 0

	for _, line := range lines {
		assignments := strings.Split(line, ",")
		fullOverlaps += checkAssignments(assignments)
		slightOverlaps += checkSlightAssignments(assignments)
	}
	fmt.Printf("Overlapping assignments: %d\n", fullOverlaps)
	fmt.Printf("Slightly overlapping assignments: %d\n", slightOverlaps)
}
