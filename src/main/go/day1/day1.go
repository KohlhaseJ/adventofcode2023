package day1

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

func Solve() {
	input, error := os.Open("../resources/input_day1.txt")
	if error != nil {
		panic(error)
	}
	defer input.Close()

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

		i, err := strconv.Atoi(concat)
		if err != nil {
			panic(error)
		}
		number += i
	}

	fmt.Println("Sum of all numbers:", number)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
