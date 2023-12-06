package main

import (
	"advent2023-go/day01"
	"advent2023-go/day02"
	"advent2023-go/day03"
	"advent2023-go/day04"
	"advent2023-go/day05"
	"advent2023-go/helpers"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Define an interface instantiated by each day
type Days interface {
	Run(io.Reader, int) string
}

// Run against day of your choosing, default all
func main() {
	args := os.Args[1:]
	var day int
	var err error
	if len(args) >= 1 {
		day, err = strconv.Atoi(args[0])
		if err != nil || day > 25 || day < 1 {
			usage()
		}
	}
	fmt.Printf("Running for day%02d\n", day)
	if day > 0 {
		runDay(getDay(day), day, helpers.PartA)
	} else {
		runDay(getDay(1), 1, helpers.PartA)
	}
}

func usage() {
	fmt.Printf("Usage:\n\t%v [day]\n", os.Args[0])
	os.Exit(1)
}

func getDay(day int) Days {
	switch day {
	case 2:
		return day02.Day02{}
	case 3:
		return day03.Day03{}
	case 4:
		return day04.Day04{}
	case 5:
		return day05.Day05{}
	case 6:
		//return day06.Day06{}
	case 13:
		//return day13.Day13{}
	default:

	}
	return day01.Day01{}
}

func runDay(d Days, day, part int) {
	filename := fmt.Sprintf("problems/input%02d.dat", day)
	data, err := os.ReadFile(filename)
	if err == nil {
		answerA := d.Run(bytes.NewReader(data), helpers.PartA)
		answerB := d.Run(bytes.NewReader(data), helpers.PartB)
		fmt.Printf("Day %2d answers are %v and %v\n", day, answerA, answerB)
	}
}
