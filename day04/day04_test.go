package day04

import (
	"advent2023-go/helpers"
	"reflect"
	"strings"
	"testing"
)

const example string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "13"
		got := Day04{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "0"
		got := Day04{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestNewCard(t *testing.T) {
	expect := Card{id: 1, have: []int{1}, winning: []int{1,20}, points: 1}
	got := NewCard("Card  1:  1 |  1 20")
	if ! reflect.DeepEqual(expect, got) {
		t.Errorf("Expected %v, got %v", expect, got) 
	}
}

func TestScoreCard(t *testing.T) {
	data := Card{id: 1, have: []int{1}, winning: []int{1,20}}
	got := scoreCard(data)
	expect := 1
	helpers.Check(t, expect, got)
}

func TestContains(t *testing.T) {
	a := []int{1, 2, 3, 4}
	v := 3
	if ! contains(a,v) {
		t.Errorf("contains function could not find %v in %v", v, a)
	}
}