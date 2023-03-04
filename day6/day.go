package day6

import (
	"advent_of_code/myutils"
	"fmt"
)

func CheckUnique(window []rune) bool {
	set := make(myutils.Set, len(window))
	for _, char := range window {
		if char == 0 {
			return false
		}
		if set.Contains(string(char)) {
			return false
		}
		set[string(char)] = 1
	}
	return true
}

func Day6() {
	lines := myutils.ReadFileLines("./day6/input")
	MARKER_LENGTH := 14

	for _, line := range lines {
		window := make([]rune, MARKER_LENGTH)
		for i, char := range line {
			window[i%len(window)] = char
			if CheckUnique(window) {
				fmt.Printf("marker (%c) found at %d.\n", window, i+1)
				return
			}
		}
	}
}
