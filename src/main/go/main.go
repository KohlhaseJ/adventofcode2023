package main

import (
	"flag"
	"log"
	"os"
	"slices"
	"strconv"

	"adventofcode.com/main/puzzles"
)

func readInputFile(path string, day int) *os.File {
	input, error := os.Open(path + "/day" + strconv.Itoa(day) + ".txt")
	if error != nil {
		panic(error)
	}

	return input
}

func main() {
	var path string
	flag.StringVar(&path, "basePath", "../input", "the root path for puzzle input files")
	var day int
	flag.IntVar(&day, "day", 1, "the day of the puzzle to solve")
	flag.Parse()

	days := []int{1, 2, 3}
	if slices.Contains(days, day) == false {
		log.Fatal("not expected day: ", day)
	}

	input := readInputFile(path, day)
	defer input.Close()

	switch day {
	case 1:
		puzzles.SolveDay1(input)
	case 2:
		puzzles.SolveDay2(input)
	case 3:
		puzzles.SolveDay3(input)
	}
}
