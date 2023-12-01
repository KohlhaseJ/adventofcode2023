package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"adventofcode.com/main/puzzles"
)

func readInputFile(path string, day int) *os.File {
	input, error := os.Open(path + "/day" + strconv.Itoa(day) + ".txt")
	if error != nil {
		panic(error)
	}
	defer input.Close()

	return input
}

func main() {
	var path string
	flag.StringVar(&path, "basePath", "../input", "The root path for input files. The files in that path are expected to be named day1.txt, day2.txt, ...")
	var day int
	flag.IntVar(&day, "day", 1, "The day of the puzzle to solve")
	flag.Parse()

	switch day {
	case 1:
		puzzles.SolveDay1(readInputFile(path, day))
	default:
		log.Fatal("Not expected day: ", day)
	}
}
