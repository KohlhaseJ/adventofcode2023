package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func digits() [9]string {
	return [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

func digitsAsText() [9]string {
	return [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
}

func SolveDay1(input *os.File) {
	scanner := bufio.NewScanner(input)
	regex := regexp.MustCompile(`\d`)
	number := 0
	for scanner.Scan() {
		textLine := scanner.Text()
		for index, digitAsText := range digitsAsText() {
			textLine = strings.ReplaceAll(textLine, digitAsText, digitAsText+digits()[index]+digitAsText)
		}
		digits := regex.FindAllString(textLine, -1)
		concat := digits[0] + digits[len(digits)-1]

		intValue, error := strconv.Atoi(concat)
		if error != nil {
			panic(error)
		}
		number += intValue
	}

	if error := scanner.Err(); error != nil {
		panic(error)
	}

	fmt.Println("Sum of calibration values:", number)
}
