package day01

import (
	"advent2023-go/helpers"
	"strings"
	"testing"
)

const exampleA string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const exampleB string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "142"
		got := Day01{}.Run(strings.NewReader(exampleA), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "281"
		got := Day01{}.Run(strings.NewReader(exampleB), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestGetNumber(t *testing.T) {
	expect := 12
	got := getNumber([]byte(`x1a7bc2ff`))
	helpers.Check(t, expect, got)
}

func TestGetParsedDigits(t *testing.T) {
	t.Run("Test parsed digits function", func(t *testing.T) {
		test := []byte(`twoneighthreeseveninefour4fiveight`)
		expect := 28
		got := getParsedDigits(test)
		helpers.Check(t, expect, got)
	})
}
func TestDigitParserHelpers(t *testing.T) {
	t.Run("Test written digit parser", func(t *testing.T) {
		expect := 7
		got := parseDigit("seven")
		helpers.Check(t, expect, got)
	})
	t.Run("Test numeric digit parser", func(t *testing.T) {
		expect := 7
		got := parseDigit("7")
		helpers.Check(t, expect, got)
	})
	t.Run("Test string reverser", func(t *testing.T) {
		expect := "gnilleps"
		got := string(reverse([]byte("spelling")))
		helpers.CheckString(t, expect, got)
	})
}
