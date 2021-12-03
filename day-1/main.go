package main

import (
	"fmt"
	"strconv"

	"github.com/garciademarina/adventofcode/common"
)

func main() {
	result := SolvePart1("/input.txt")
	fmt.Printf("Result Part 1: %d\n", result)

	result = SolvePart2("/input.txt")
	fmt.Printf("Result Part 2: %d\n", result)
}

// SolvePart1 ...
func SolvePart1(path string) int {

	var counter int
	var lastMeasurement *int

	for _, line := range common.ReadFile(path) {
		currentMeasurement, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if lastMeasurement != nil && currentMeasurement > *lastMeasurement {
			counter++
		}
		lastMeasurement = &currentMeasurement
	}

	return counter
}

// SolvePart2 ...
func SolvePart2(path string) int {

	// 199  A			index = 0
	// 200  A B			index = 1
	// 208  A B C		index = 2
	// 210    B C D		index = 3
	// 200  E   C D		index = 4
	// 207  E F   D		index = 5
	// 240  E F G		index = 6	6+3 = 9
	// 269    F G H		index = 7	7+3 = 10
	// 260      G H		index = 8
	// 263        H		index = 9
	var counter int
	var lastMeasurement *int

	lines := common.ReadFile(path)
	totalLines := len(lines)

	if totalLines <= 3 {
		return 0
	}
	for index, line := range lines {
		currentValue1, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentValue2, err := strconv.Atoi(lines[index+1])
		if err != nil {
			panic(err)
		}
		currentValue3, err := strconv.Atoi(lines[index+2])
		if err != nil {
			panic(err)
		}

		currentMeasurement := currentValue1 * currentValue2 * currentValue3

		if lastMeasurement != nil && currentMeasurement > *lastMeasurement {
			counter++
		}
		lastMeasurement = &currentMeasurement

		if index+3 > totalLines-1 {
			return counter
		}
	}

	return counter
}
