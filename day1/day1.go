package day1

import (
	"advent_of_code/myutils"
	"fmt"
	"sort"
	"strconv"
)

func Day1() {
	lines := myutils.ReadFileLines("./day1/input")

	top3 := make([]int, 3)
	var summed int
	for _, line := range lines {
		if line == "" {
			// fmt.Println(summed)
			top3 = append(top3, summed)
			sort.Ints(top3)
			top3 = top3[1:]
			summed = 0
		} else {
			sumPart, err := strconv.Atoi(line)
			myutils.Check(err)
			summed += sumPart
		}
	}
	fmt.Printf("Top 3 calories count: %v\n", top3)
	fmt.Printf("Sum of top 3 calories count: %d\n", myutils.Sum(top3))
}
