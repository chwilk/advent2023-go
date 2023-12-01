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
		answer += getParsedDigits(scanner.Bytes())
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

func getParsedDigits(line []byte) (answer int) {
	first := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`).FindIndex(line)
	rev := reverse(line)
	tmp := regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`).FindIndex(rev)
	answer = 10 * parseDigit(string(line[first[0]:first[1]]))

	// unreverse indices for last
	var last [2]int
	length := len(line)
	last[0] = length - tmp[1]
	last[1] = length - tmp[0]
	answer += parseDigit(string(line[last[0]:last[1]]))

	return
}

func parseDigit(text string) (answer int) {
	switch text {
	case "one":
		answer = 1
	case "two":
		answer = 2
	case "three":
		answer = 3
	case "four":
		answer = 4
	case "five":
		answer = 5
	case "six":
		answer = 6
	case "seven":
		answer = 7
	case "eight":
		answer = 8
	case "nine":
		answer = 9
	default:
		var err error
		answer, err = strconv.Atoi(text)
		if err != nil {
			panic("Failed to parse a number")
		}
	}
	return
}

// reverse a byte array
func reverse(a []byte) (b []byte) {
	b = make([]byte, len(a))
	copy(b, a)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
