package main

import (
	"fmt"
	"strconv"
	"strings"

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
	lines := common.ReadFile(path)

	// how many digits we have at the first line?
	digits := strings.Split(lines[0], "")
	numDigits := len(digits)

	mapCounter := make(map[int][]int, numDigits)
	for _, line := range lines {
		digits := strings.Split(line, "")
		for i, value := range digits {
			aux, _ := strconv.Atoi(value)
			mapCounter[i] = append(mapCounter[i], aux)
		}

	}

	numWhenMoreTimes := ""
	numWhenLessTimes := ""
	for i := 0; i < len(mapCounter); i++ {
		v := mapCounter[i]
		one := 0
		zero := 0

		for _, aux := range v {
			if aux == 0 {
				zero++
			} else {
				one++
			}
		}

		if zero > one {
			numWhenMoreTimes = fmt.Sprintf(numWhenMoreTimes+"%d", 0)
			numWhenLessTimes = fmt.Sprintf(numWhenLessTimes+"%d", 1)
		} else {
			numWhenMoreTimes = fmt.Sprintf(numWhenMoreTimes+"%d", 1)
			numWhenLessTimes = fmt.Sprintf(numWhenLessTimes+"%d", 0)
		}

	}

	return binaryStringToInteger(numWhenMoreTimes) * binaryStringToInteger(numWhenLessTimes)
}

// SolvePart2 ...
func SolvePart2(path string) int {
	oxygenGeneratorRating := 0
	CO2ScrubberRating := 0

	lines := common.ReadFile(path)
	currentOxygenList := lines
	currentCO2List := lines

	digits := strings.Split(lines[0], "")
	numDigits := len(digits)

	for index := 0; index < numDigits; index++ {
		// get oxygenGeneratorRating
		if len(currentOxygenList) > 1 {
			currentOxygenList = getNewList(index, currentOxygenList, "MORE")
			if len(currentOxygenList) == 1 {
				oxygenGeneratorRating = binaryStringToInteger(currentOxygenList[0])
			}
		}
		// get CO2ScrubberRating
		if len(currentCO2List) > 1 {
			currentCO2List = getNewList(index, currentCO2List, "LESS")
			if len(currentCO2List) == 1 {
				CO2ScrubberRating = binaryStringToInteger(currentCO2List[0])
			}
		}

	}
	return oxygenGeneratorRating * CO2ScrubberRating
}

func getNewList(index int, list []string, typeComp string) []string {
	one := 0
	zero := 0
	for _, line := range list {
		digits := strings.Split(line, "")
		aux, _ := strconv.Atoi(digits[index])
		if aux == 0 {
			zero++
		} else {
			one++
		}
	}

	var newList []string
	for _, line := range list {
		digits := strings.Split(line, "")
		aux, _ := strconv.Atoi(digits[index])
		if typeComp == "MORE" {
			if zero > one && aux == 0 {
				newList = append(newList, line)
			}
			if one > zero && aux == 1 {
				newList = append(newList, line)
			}
			if one == zero && aux == 1 {
				newList = append(newList, line)
			}
		}
		if typeComp == "LESS" {
			if zero > one && aux == 1 {
				newList = append(newList, line)
			}
			if one > zero && aux == 0 {
				newList = append(newList, line)
			}
			if one == zero && aux == 0 {
				newList = append(newList, line)
			}
		}
	}

	return newList
}

func binaryStringToInteger(input string) int {
	output, _ := strconv.ParseInt(input, 2, 64)
	return int(output)
}

func mapToString(input map[int][]int) string {
	output := ""
	for _, v := range input {
		output = fmt.Sprintf(output+"%d", v[0])
	}
	return output
}

func intToString(input []int) string {
	b := ""
	for _, v := range input {
		if len(b) > 0 {
			b += ","
		}
		b += strconv.Itoa(v)
	}

	return b
}
