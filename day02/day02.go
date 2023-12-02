package day02

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day02 struct {
}

func (f Day02) Run(input io.Reader, part int) (result string) {
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
	maximums := [4]int{0,12,13,14}
	for scanner.Scan() {
		tmp := getMax(scanner.Text())
		possible := true
		for i:=1; i<4; i++ {
			if tmp[i]>maximums[i] {
				possible = false
			}
		}
		if possible {
			answer += tmp[0]
		}
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		tmp := getMax(scanner.Text())
		power := tmp[1] * tmp[2] * tmp[3]
		answer += power
	}
	return
}

func getMax(game string) (answer [4]int) {
	var err error
	a := strings.Split(game, ":")
	g := strings.Split(a[0], " ")
	answer[0], err = strconv.Atoi(g[1])
	if err != nil {
		panic ("Could not tell the game number in getMax")
	}
	rounds := strings.Split(a[1], ";")
	for _, round := range rounds {
		for _, clr := range strings.Split(strings.Trim(round, " "), ", ") {
			pair := strings.Split(clr, " ")
			var idx int
			switch pair[1] {
			case "red":
				idx = 1
			case "green":
				idx = 2
			case "blue":
				idx = 3
			}
			num, e := strconv.Atoi(pair[0])
			if e != nil {
				panic ("Could not get a number to associate with color")
			}
			if num > answer[idx] {
				answer[idx] = num
			}
		}
	}	
	return
}