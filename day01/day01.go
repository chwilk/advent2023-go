package day01

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Day01 struct {
}

func (f Day01) Run(input io.Reader, part int) (result string) {
	switch part {
	case helpers.PartA:
		result = fmt.Sprint(PartA(input))
	default:
		result = fmt.Sprint(PartB(input))
	}
	return
}

func PartA(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		answer += getNumber(scanner.Bytes())
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		answer += getNumber(parseDigits(scanner.Bytes()))
	}
	return
}

func getNumber(line []byte) (answer int) {
	numbers := regexp.MustCompile(`\d`).FindAll(line, -1)
	tmp, err := strconv.Atoi(string(numbers[0]))
	if err == nil {
		answer = tmp * 10
	}
	tmp, err = strconv.Atoi(string(numbers[len(numbers)-1]))
	if err == nil {
		answer += tmp
	}
	return
}

func parseDigits(line []byte) (answer []byte) {
	answer = line

	for {
		var replace string
		index := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`).FindIndex(answer)
		if index == nil {
			break
		}
		switch string(answer[index[0]:index[1]]) {
		case "one":
			replace = "1"
		case "two":
			replace = "2"
		case "three":
			replace = "3"
		case "four":
			replace = "4"
		case "five":
			replace = "5"
		case "six":
			replace = "6"
		case "seven":
			replace = "7"
		case "eight":
			replace = "8"
		case "nine":
			replace = "9"
		default:
			replace = "X"
		}
		var head, tail string
		if index[0] == 0 {
			head = ""
		} else {
			head = string(answer[0:index[0]])
		}
		if index[1] == len(answer) {
			tail = ""
		} else {
			tail = string(answer[index[1]:])
		}
		answer = []byte(fmt.Sprintf("%s%s%s", head, replace, tail))
	}
	return
}
