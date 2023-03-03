// Package myutils provides a variety of useful functions
package myutils

import (
	"bufio"
	"os"
)

// Function Check checks if the error is nil, if not it panics with the error message
func Check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Function ReadFileLines reads a file and returns a slice of all lines in the file
func ReadFileLines(fileName string) []string {
	f, err := os.Open(fileName)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Function Sum sums a slice of integers and returns the results
func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

type Set map[string]int

func (s Set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}
