package main

import (
	"advent_of_code/day1"
	"advent_of_code/day2"
	"advent_of_code/day3"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <day_to_run>\n", os.Args[0])
		return
	}
	switch os.Args[1] {
	case "1":
		day1.Day1()
	case "2":
		day2.Day2()
	case "3":
		day3.Day3()
	default:
		fmt.Printf("The argument '%v' is invalid or the day was not yet implemented :)\n", os.Args[1])
	}
}
