package day03

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
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

func PartA(lines io.Reader) (answer int64) {
	schematic := readPuzzle(lines)

	re := regexp.MustCompile(`\d+`)
	for i, row := range schematic {
		matches := re.FindAllIndex(row, -1)
		for _, m := range matches {
			tmp, err := strconv.Atoi(string(row[m[0]:m[1]]))
			if err != nil {
				panic(err)
			}
			if checkAdjacent(schematic, i, m) {
				answer +=  int64(tmp)
			}
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

func readPuzzle(lines io.Reader) (puzzle [][]byte) {
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line), 140)
		for i, v := range line {
			switch {
			case v == '.':
				// noop
			case v > 47 && v < 58:
				row[i] = line[i]
			default:
				row[i] = 'S'
			}
		}
		puzzle = append(puzzle, row)
	}

	return
}

func checkAdjacent(data [][]byte, row int, cols []int) bool {
	lastRow := len(data) - 1
	maxCols := len(data[0])

	checkRange := [2][2]int{{0,0},{0,0}}
	if row == 0 {
		checkRange[0][1] = 0
		checkRange[1][1] = 1
	} else if row == lastRow {
		checkRange[0][1] = row -1
		checkRange[1][1] = lastRow
	} else {
		checkRange[0][1] = row - 1
		checkRange[1][1] = row + 1
	}
	if cols[0] == 0 {
		checkRange[0][0] = 0
		checkRange[1][0] = cols[1]
	} else if cols[1] == maxCols {
		checkRange[0][0] = cols[0] - 1
		checkRange[1][0] = maxCols - 1
	} else {
		checkRange[0][0] = cols[0] - 1
		checkRange[1][0] = cols[1]
	}

	for i := checkRange[0][1]; i <= checkRange[1][1]; i++ {
		if strings.Contains(string(data[i][checkRange[0][0]:checkRange[1][0]+1]), "S") {
			return true
		}
	}

	return false
}
