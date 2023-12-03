package day03

import (
	"advent2023-go/helpers"
	"reflect"
	"strings"
	"testing"
)

const example string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "4361"
		got := Day03{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "0"
		got := Day03{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestReadPuzzle(t *testing.T) {
	expect := [][]byte{
		{'4', '6', '7', 0, 0, '1', '1', '4', 0, 0},
		{0, 0, 0, 'S', 0, 0, 0, 0, 0, 0},
		{0, 0, '3', '5', 0, 0, '6', '3', '3', 0},
		{0, 0, 0, 0, 0, 0, 'S', 0, 0, 0},
		{'6', '1', '7', 'S', 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 'S', 0, '5', '8', 0},
		{0, 0, '5', '9', '2', 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, '7', '5', '5', 0},
		{0, 0, 0, 'S', 0, 'S', 0, 0, 0, 0},
		{0, '6', '6', '4', 0, '5', '9', '8', 0, 0}}
	got := readPuzzle(strings.NewReader(example))
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestCheckAdjacent(t *testing.T) {
	t.Run("Easy case, adjacent symbol", func(t *testing.T) {
		data := [][]byte{
			{0, 0, 0, 0, 0},
			{0, '4', '4', '4', 0},
			{0, 0, 0, 0, 'S'}}
		got := checkAdjacent(data, 1, []int{1, 4})
		if got == false {
			t.Error("Missed adjacent symbol")
		}
	})
	t.Run("Edge case, top", func(t *testing.T) {
		data := [][]byte{
			{0, '4', '4', '4', 0},
			{0, 0, 0, 0, 'S'}}
		got := checkAdjacent(data, 0, []int{1, 4})
		if got == false {
			t.Error("Missed adjacent symbol")
		}
	})
	t.Run("Edge case, bottom", func(t *testing.T) {
		data := [][]byte{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 'S'},
			{0, '4', '4', '4', 0}}
		got := checkAdjacent(data, 2, []int{1, 4})
		if got == false {
			t.Error("Missed adjacent symbol")
		}
	})
	t.Run("Edge case, bottom left", func(t *testing.T) {
		data := [][]byte{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 'S', 0},
			{'4', '4', '4', 0, 0}}
		got := checkAdjacent(data, 2, []int{0, 3})
		if got == false {
			t.Error("Missed adjacent symbol")
		}
	})
	t.Run("Edge case, top right", func(t *testing.T) {
		data := [][]byte{
			{0,0,'4', '4', '4'},
			{0, 'S', 0, 0, 0},
			{0, 0, 0, 0, 0}}
		got := checkAdjacent(data, 0, []int{2, 5})
		if got == false {
			t.Error("Missed adjacent symbol")
		}
	})
}
