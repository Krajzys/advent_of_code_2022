package day5

import (
	"advent_of_code/myutils"
	"fmt"
	"strings"
)

type Stack map[int][]rune

func Day5() {
	INITIAL_COLUMN := 1
	SPACE_BETWEEN := 4
	MOVE_MANY := false

	lines := myutils.ReadFileLines("./day5/input")
	stacksCount := (len(lines[0]) / SPACE_BETWEEN) + 1
	myStack := make(Stack, stacksCount)
	stackNo := 1

	for _, line := range lines {
		stackNo = 1
		if !strings.Contains(line, "[") {
			if strings.HasPrefix(line, "move ") {
				instruction := strings.Split(line, " ")
				qty := myutils.AtoiCheck(instruction[1])
				from := myutils.AtoiCheck(instruction[3])
				to := myutils.AtoiCheck(instruction[5])
				fmt.Printf("qty: %d, from %d, to %d\n", qty, from, to)
				if MOVE_MANY {
					fromCrates := myStack[from][len(myStack[from])-qty:]
					myStack[from] = myStack[from][:len(myStack[from])-qty]
					myStack[to] = append(myStack[to], fromCrates...)
				} else {
					for i := 0; i < qty; i++ {
						fromCrate := myStack[from][len(myStack[from])-1]
						myStack[from] = myStack[from][:len(myStack[from])-1]
						myStack[to] = append(myStack[to], fromCrate)
					}
				}
				fmt.Printf("Instruction: %s\n", line)
				topCrates := ""
				for stackNo <= stacksCount {
					fmt.Printf("%d: %c\n", stackNo, myStack[stackNo])
					if len(myStack[stackNo]) > 0 {
						topCrates += string(myStack[stackNo][len(myStack[stackNo])-1])
					} else {
						topCrates += " "
					}
					stackNo += 1
				}
				fmt.Printf("Top crates: %s\n", topCrates)
			} else {
				topCrates := ""
				for stackNo <= stacksCount {
					fmt.Printf("%d: %c\n", stackNo, myStack[stackNo])
					if len(myStack[stackNo]) > 0 {
						topCrates += string(myStack[stackNo][len(myStack[stackNo])-1])
					} else {
						topCrates += " "
					}
					stackNo += 1
				}
				fmt.Printf("Top crates: %s\n", topCrates)
			}
		} else {
			i := INITIAL_COLUMN
			for i <= len(line) {
				if line[i] != ' ' {
					tempSlice := make([]rune, 1)
					tempSlice[0] = rune(line[i])
					myStack[stackNo] = append(tempSlice, myStack[stackNo]...)
				}
				i += SPACE_BETWEEN
				stackNo += 1
			}
			if stackNo > stacksCount {
				stacksCount = stackNo - 1
			}
		}
	}
}
