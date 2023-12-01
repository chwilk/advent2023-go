package day01

import (
	"advent2022-go/helpers"
	"bufio"
	"fmt"
	"io"
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
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
	}
	return
}
