package day3

import (
	"advent_of_code/myutils"
	"fmt"
)

type runeset map[int]rune
type set map[int]string
type intset map[int]int

func runeToPoints(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return (int(char) - int('A') + 27)
	}
	if char >= 'a' && char <= 'z' {
		return (int(char) - int('a') + 1)
	}
	return 0
}

func Day3() {
	lines := myutils.ReadFileLines("./day3/input")

	repetitionSum := 0
	for _, line := range lines {
		firstComp := line[:len(line)/2]
		firstCompSet := make(set)
		secondComp := line[len(line)/2:]
		repetitions := make(set)
		for _, el := range firstComp {
			firstCompSet[runeToPoints(el)] = string(el)
		}
		for _, el := range secondComp {
			if _, ok := firstCompSet[runeToPoints(el)]; ok {
				if _, ok := repetitions[runeToPoints(el)]; !ok {
					repetitions[runeToPoints(el)] = string(el)
					repetitionSum += runeToPoints(el)
				}
			}
		}
	}
	fmt.Printf("Sum of repetition points: %d\n", repetitionSum)

	repetitionSum = 0
	repetitions := make(intset)
	for i, line := range lines {
		uniqueItems := make(runeset)
		for _, el := range line {
			uniqueItems[runeToPoints(el)] = el
		}
		for _, el := range uniqueItems {
			if _, ok := repetitions[runeToPoints(el)]; !ok {
				repetitions[runeToPoints(el)] = 1
			} else {
				repetitions[runeToPoints(el)] += 1
			}
		}
		if (i+1)%3 == 0 {
			for points, reps := range repetitions {
				if reps == 3 {
					repetitionSum += points
				}
			}
			repetitions = make(intset)
		}
	}
	fmt.Printf("Sum of repetition points (v2): %d\n", repetitionSum)
}
