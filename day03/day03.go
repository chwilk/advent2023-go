package day03

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Day03 struct {
}

func (f Day03) Run(input io.Reader, part int) (result string) {
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
	var schematic [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		schematic = append(schematic, line)
	}
	re := regexp.MustCompile(`(\d+)`)
	for _, row := range schematic {
		matches := re.FindAllIndex(row, -1)
		for _, m := range matches {
			tmp, err := strconv.Atoi(string(row[m[0]:m[1]]))
			if err != nil {
				panic (err)
			}
			answer += tmp
		}
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
	}
	return
}