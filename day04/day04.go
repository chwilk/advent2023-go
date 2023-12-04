package day04

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Day04 struct {
}

type Card struct {
	id      int
	have    []int
	winning []int
	matches  int
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
		if card.matches > 0 {
			answer += int(math.Pow(2, float64(card.matches - 1)))
		}
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)
	var pile []Card
	for scanner.Scan() {
		card:= NewCard(scanner.Text())
		pile = append(pile, card)
	}
	copies := make([]int, len(pile)+1)
	for _, card := range pile {
		copies[card.id] += 1 // count original card we're iterating through
		for j := 1; j <= card.matches; j++ {
			copies[card.id+j] += copies[card.id] 
		}
		answer += copies[card.id]
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
	card.matches = scoreCard(card)
	return card
}

func scoreCard(card Card) (score int) {
	for _, num := range card.have {
		if  contains(card.winning, num) {
			score ++
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

func sum (a []int) (answer int) {
	for _,v := range a {
		answer += v
	}
	return
}