package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func colorCountMap() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func initColorCountMap() map[string]int {
	return map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
}

func SolveDay2(input *os.File) {

	scanner := bufio.NewScanner(input)
	sumOfValidGames := 0
	power := 0

	for scanner.Scan() {
		textLine := scanner.Text()

		game := strings.Split(strings.Split(textLine, ":")[0], " ")[1]
		gameInput := strings.Split(textLine, ":")[1]

		sets := strings.Split(gameInput, ";")

		isValid := true

		colorPowerMap := initColorCountMap()

		for _, set := range sets {
			set = strings.TrimSpace(set)
			cubesPerColor := strings.Split(set, ",")

			inputColorCountMap := initColorCountMap()

			for _, cubesInColor := range cubesPerColor {
				cubesInColor = strings.TrimSpace(cubesInColor)
				count, _ := strconv.Atoi(strings.Split(cubesInColor, " ")[0])
				color := strings.Split(cubesInColor, " ")[1]
				inputColorCountMap[color] = inputColorCountMap[color] + count
				if count > colorPowerMap[color] {
					colorPowerMap[color] = count
				}
			}

			for key, value := range inputColorCountMap {
				if colorCountMap()[key] < value {
					isValid = false
				}
			}
		}

		if isValid {
			intValue, _ := strconv.Atoi(game)
			sumOfValidGames += intValue
		}

		setPower := 1
		for _, value := range colorPowerMap {
			setPower = setPower * value
		}
		power += setPower

	}

	if error := scanner.Err(); error != nil {
		panic(error)
	}

	fmt.Println("sum of valid game ids:", sumOfValidGames)
	fmt.Println("sum of powers:", power)
}
