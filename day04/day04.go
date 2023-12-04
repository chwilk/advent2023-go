package day04

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day04 struct {
}

type Card struct {
	id      int
	have    []int
	winning []int
	points  int
}

func (f Day04) Run(input io.Reader, part int) (result string) {
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
		card:= NewCard(scanner.Text())
		answer += card.points
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
	}
	return
}

func NewCard(row string) (card Card) {
	tokens := strings.Split(row, " ")
	var half bool
	for _, tok := range tokens[1:] {
		switch tok {
		case "|":
			half = true
		case "":
			// noop if double spaces
		default:
			var err error
			length := len(tok)
			if tok[length-1] == ':' {
				card.id, err = strconv.Atoi(tok[0:length-1])
				if err != nil {
					panic(err)
				}
			} else {
				var tmp int
				tmp, err = strconv.Atoi(tok)
				if err == nil {
					switch half {
					case false:
						card.have = append(card.have, tmp)
					default:
						card.winning = append(card.winning, tmp)
					}
				}
			}
		}
	}
	card.points = scoreCard(card)
	return card
}

func scoreCard(card Card) (score int) {
	for _, num := range card.have {
		if  contains(card.winning, num) {
			if score > 0 {
				score *= 2
			} else {
				score = 1
			}
		}
	}
	return
}

func contains(a []int, val int) bool {
	for _, v := range a {
		if v == val {
			return true
		}
	}
	return false
}