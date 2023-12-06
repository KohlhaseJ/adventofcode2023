package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func buildNumberWindow(lineWindow []string) [][]int {
	numberRegex := regexp.MustCompile("[0-9]+")
	numberWindow := make([][]int, len(lineWindow))

	for lineIndex := 0; lineIndex < len(lineWindow); lineIndex++ {
		numberWindow[lineIndex] = make([]int, len(lineWindow[lineIndex]))
		numberMatches := numberRegex.FindAllString(lineWindow[lineIndex], -1)
		numberMatchIndices := numberRegex.FindAllStringIndex(lineWindow[lineIndex], -1)

		for matchIndex, numberMatchIndexRange := range numberMatchIndices {
			number, _ := strconv.Atoi(numberMatches[matchIndex])
			for rangeIndex := numberMatchIndexRange[0]; rangeIndex < numberMatchIndexRange[1]; rangeIndex++ {
				numberWindow[lineIndex][rangeIndex] = number
			}
		}
	}

	return numberWindow
}

func findAdjacentNumbers(index int, numberWindow [][]int) []int {
	adjacents := []int{}

	if numberWindow[0][index] > 0 {
		adjacents = append(adjacents, numberWindow[0][index])
	} else {
		if index > 0 && numberWindow[0][index-1] > 0 {
			adjacents = append(adjacents, numberWindow[0][index-1])
		}

		if index < len(numberWindow[0]) && numberWindow[0][index+1] > 0 {
			adjacents = append(adjacents, numberWindow[0][index+1])
		}
	}

	if index > 0 && numberWindow[1][index-1] > 0 {
		adjacents = append(adjacents, numberWindow[1][index-1])
	}

	if index < len(numberWindow[1]) && numberWindow[1][index+1] > 0 {
		adjacents = append(adjacents, numberWindow[1][index+1])
	}

	if numberWindow[2][index] > 0 {
		adjacents = append(adjacents, numberWindow[2][index])
	} else {
		if index > 0 && numberWindow[2][index-1] > 0 {
			adjacents = append(adjacents, numberWindow[2][index-1])
		}

		if index < len(numberWindow[2]) && numberWindow[2][index+1] > 0 {
			adjacents = append(adjacents, numberWindow[2][index+1])
		}
	}

	return adjacents
}

func SolveDay3(input *os.File) {

	symbolRegex := regexp.MustCompile(`[^\d.]`)
	starRegex := regexp.MustCompile("[*]")
	scanner := bufio.NewScanner(input)

	var lineWindow []string

	partNumbersSum := 0
	gearRatioSum := 0

	for scanner.Scan() {
		textLine := scanner.Text()

		lineWindow = append(lineWindow, textLine)
		if len(lineWindow) > 3 {
			lineWindow = lineWindow[1:]
		}
		if len(lineWindow) < 3 {
			continue
		}

		numberWindow := buildNumberWindow(lineWindow)

		symbolIndices := symbolRegex.FindAllStringIndex(lineWindow[1], -1)
		if len(symbolIndices) > 0 {
			for _, indexRange := range symbolIndices {
				index := indexRange[0]

				adjacents := findAdjacentNumbers(index, numberWindow)
				for _, number := range adjacents {
					partNumbersSum += number
				}
			}
		}

		starIndices := starRegex.FindAllStringIndex(lineWindow[1], -1)
		if len(starIndices) > 0 {
			for _, indexRange := range starIndices {
				index := indexRange[0]
				adjacents := findAdjacentNumbers(index, numberWindow)

				if len(adjacents) == 2 {
					gearRatioSum += (adjacents[0] * adjacents[1])
				}
			}
		}

	}

	if error := scanner.Err(); error != nil {
		panic(error)
	}

	fmt.Println("Sum of part numbers", partNumbersSum)
	fmt.Println("Sum of gear ratios", gearRatioSum)
}
